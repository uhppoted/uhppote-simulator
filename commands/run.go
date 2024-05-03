package commands

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"os/signal"
	"regexp"
	"time"

	codec "github.com/uhppoted/uhppote-core/encoding/UTO311-L0x"
	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/log"
	"github.com/uhppoted/uhppote-simulator/rest"
	"github.com/uhppoted/uhppote-simulator/simulator"
)

var debug bool = false

func Simulate(ctx *simulator.Context, dbg bool) {
	debug = dbg

	if bind, err := net.ResolveUDPAddr("udp4", ctx.BindAddress); err != nil {
		log.Errorf("failed to resolve UDP bind address [%v]", err)
	} else if udp, err := net.ListenUDP("udp", bind); err != nil {
		log.Errorf("failed to bind to UDP socket [%v]", err)
	} else {
		defer udp.Close()

		log.Infof("bound to address %s", bind)

		wait := make(chan int, 1)
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		go run(ctx, udp, wait)

		<-interrupt
		udp.Close()
		<-wait

		os.Exit(1)
	}
}

func run(ctx *simulator.Context, udp *net.UDPConn, wait chan int) {
	bind, err := net.ResolveUDPAddr("udp4", ctx.BindAddress)
	if err != nil {
		log.Errorf("failed to resolve UDP bind address [%v]", err)
		return
	}

	go func() {
		if err := udpListenAndServe(ctx, udp); err != nil {
			fmt.Printf("%v\n", err)
		}
		wait <- 0
	}()

	go func() {
		for {
			msg := ctx.DeviceList.GetMessage()

			if msg.Event {
				sendto(bind, msg.Destination, msg.Message)
			} else {
				send(udp, msg.Destination, msg.Message)
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		last := "00:00"
		for {
			<-ticker.C
			now := time.Now().Format("15:34")
			if now != last {
				last = now
				tasks(ctx)
			}
		}
	}()

	go func() {
		rest.Run(ctx)
	}()

}

func udpListenAndServe(ctx *simulator.Context, c *net.UDPConn) error {
	for {
		if request, remote, err := receive(c); err != nil {
			return err
		} else {
			handle(ctx, c, remote, request)
		}
	}
}

func handle(ctx *simulator.Context, c *net.UDPConn, src *net.UDPAddr, bytes []byte) {
	if request, err := messages.UnmarshalRequest(bytes); err != nil {
		log.Errorf("%v", err)
	} else {
		f := func(s simulator.Simulator) {
			if response, err := s.Handle(request); err != nil {
				warnf(tag(request), "%v", err)
			} else if response != nil {
				s.Send(src, response)
			}
		}

		ctx.DeviceList.Apply(f)
	}
}

func tasks(ctx *simulator.Context) {
	f := func(s simulator.Simulator) {
		s.RunTasks()
	}

	ctx.DeviceList.Apply(f)
}

func receive(c *net.UDPConn) ([]byte, *net.UDPAddr, error) {
	request := make([]byte, 2048)

	N, remote, err := c.ReadFromUDP(request)
	if err != nil {
		return []byte{}, nil, fmt.Errorf("failed to read from UDP socket [%v]", err)
	} else if debug {
		log.Debugf("received %v bytes from %v\n%s", N, remote, dump(request[0:N], " ...          "))
	}

	return request[:N], remote, nil
}

func send(c *net.UDPConn, dest *net.UDPAddr, message any) {
	msg, err := codec.Marshal(message)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	N, err := c.WriteToUDP(msg, dest)
	if err != nil {
		log.Errorf("failed to write to UDP socket [%v]", err)
	} else if debug {
		log.Infof("sent %v bytes to %v\n%s", N, dest, dump(msg[0:N], " ...          "))
	}
}

func sendto(bind *net.UDPAddr, dest *net.UDPAddr, message any) {
	var addr *net.UDPAddr

	if bind != nil && bind.IP != nil {
		addr = &net.UDPAddr{
			IP:   bind.IP.To4(),
			Port: 0,
			Zone: bind.Zone,
		}
	}

	msg, err := codec.Marshal(message)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	if c, err := net.DialUDP("udp4", addr, dest); err != nil {
		log.Errorf("failed to create UDP event socket [%v]", err)
	} else {
		defer c.Close()

		N, err := c.Write(msg)
		if err != nil {
			log.Errorf("failed to write to UDP socket [%v]", err)
		} else if debug {
			log.Infof("sent %v bytes to %v\n%s", N, dest, dump(msg[0:N], " ...          "))
		}
	}
}

func warnf(tag any, format string, args ...any) {
	f := fmt.Sprintf("%-10v  %v", tag, format)

	log.Warnf(f, args...)
}

func dump(m []byte, prefix string) string {
	return regexp.MustCompile("(?m)^(.*)").ReplaceAllString(hex.Dump(m), prefix+"$1")
}

func tag(rq any) string {
	switch rq.(type) {
	case *messages.ActivateAccessKeypadsRequest:
		return "activate-keypads"

	case *messages.AddTaskRequest:
		return "add-task"

	case *messages.ClearTaskListRequest:
		return "clear-tasklist"

	case *messages.ClearTimeProfilesRequest:
		return "clear-profiles"

	case *messages.DeleteCardRequest:
		return "delete-card"

	case *messages.DeleteCardsRequest:
		return "delete-cards"

	case *messages.GetCardByIDRequest:
		return "get-card"

	case *messages.GetCardByIndexRequest:
		return "get-card-by-index"

	case *messages.GetDeviceRequest:
		return "get-device"

	case *messages.GetDoorControlStateRequest:
		return "get-door-control"

	case *messages.GetEventRequest:
		return "get-event"

	case *messages.GetEventIndexRequest:
		return "get-event-index"

	case *messages.GetListenerRequest:
		return "get-listener"

	case *messages.GetStatusRequest:
		return "get-status"

	case *messages.GetTimeRequest:
		return "get-time"

	case *messages.GetTimeProfileRequest:
		return "get-time-profile"

	case *messages.OpenDoorRequest:
		return "open-door"

	case *messages.PutCardRequest:
		return "put-card"

	case *messages.GetCardsRequest:
		return "get-cards"

	case *messages.SetDoorControlStateRequest:
		return "set-door-control"

	case *messages.SetDoorPasscodesRequest:
		return "set-passcodes"

	case *messages.SetListenerRequest:
		return "set-listener"

	case *messages.SetAddressRequest:
		return "set-address"

	case *messages.SetEventIndexRequest:
		return "set-event-index"

	case *messages.RecordSpecialEventsRequest:
		return "record-special-events"

	case *messages.SetTimeProfileRequest:
		return "set-time-profile"

	case *messages.RefreshTaskListRequest:
		return "refresh-tasklist"

	case *messages.SetPCControlRequest:
		return "set-pc-control"

	case *messages.SetInterlockRequest:
		return "set-interlock"

	case *messages.SetTimeRequest:
		return "set-time"

	case *messages.RestoreDefaultParametersRequest:
		return "restore-default-parameters"

	default:
		return "???"
	}
}
