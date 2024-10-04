package commands

import (
	"fmt"
	"net"
	"net/netip"
	"os"
	"os/signal"
	"reflect"
	"time"

	codec "github.com/uhppoted/uhppote-core/encoding/UTO311-L0x"
	"github.com/uhppoted/uhppote-core/messages"

	"github.com/uhppoted/uhppote-simulator/log"
	"github.com/uhppoted/uhppote-simulator/rest"
	"github.com/uhppoted/uhppote-simulator/simulator"
	"github.com/uhppoted/uhppote-simulator/simulator/UT0311L04"
)

var debug bool = false

func Simulate(ctx *simulator.Context, dbg bool) {
	debug = dbg

	var udp *net.UDPConn
	var tcp *net.TCPListener

	{
		bind, err := net.ResolveUDPAddr("udp4", ctx.BindAddress)
		if err != nil {
			log.Errorf("failed to resolve UDP bind address [%v]", err)
			return
		}

		udp, err = net.ListenUDP("udp", bind)
		if err != nil {
			log.Errorf("failed to bind to UDP socket [%v]", err)
			return
		}

		infof("udp", "bound to address %s", bind)
	}

	defer udp.Close()

	{
		bind, err := net.ResolveTCPAddr("tcp4", ctx.BindAddress)
		if err != nil {
			log.Errorf("failed to resolve TCP bind address [%v]", err)
			return
		}
		tcp, err = net.ListenTCP("tcp", bind)
		if err != nil {
			log.Errorf("failed to bind to TCP socket [%v]", err)
			return
		}

		infof("tcp", "bound to address %s", bind)
	}

	defer tcp.Close()

	wait := make(chan int, 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go run(ctx, udp, tcp, wait)

	<-interrupt
	udp.Close()
	<-wait

	os.Exit(1)
}

func run(ctx *simulator.Context, udp *net.UDPConn, tcp *net.TCPListener, wait chan int) {
	bind, err := net.ResolveUDPAddr("udp4", ctx.BindAddress)
	if err != nil {
		log.Errorf("failed to resolve UDP bind address [%v]", err)
		return
	}

	g := func(dest netip.AddrPort, event any) {
		if dest.IsValid() {
			sendto(bind, net.UDPAddrFromAddrPort(dest), event)
		}
	}

	UT0311L04.SetOnEvent(g)

	go func() {
		if err := udpListenAndServe(ctx, udp); err != nil {
			errorf("udp", "%v", err)
		}
		wait <- 0
	}()

	go func() {
		if err := tcpListenAndServe(ctx, tcp); err != nil {
			errorf("tcp", "%v", err)
		}
		wait <- 0
	}()

	// ... auto-send tick
	go func() {
		ticker := time.Tick(100 * time.Millisecond)

		for {
			<-ticker
			tick(ctx)
		}
	}()

	// ... on-the-minute 'tick'
	go func() {
		ticker := time.Tick(1 * time.Second)

		last := "00:00"
		for {
			<-ticker
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

func udpListenAndServe(ctx *simulator.Context, udp *net.UDPConn) error {
	handle := func(addr *net.UDPAddr, bytes []byte) {
		if request, err := messages.UnmarshalRequest(bytes); err != nil {
			errorf("udp", "%v", err)
		} else {
			f := func(s simulator.Simulator) {
				if response, err := s.Handle(request); err != nil {
					warnf(tag(request), "%v", err)
				} else if !isNil(response) {
					if msg, err := codec.Marshal(response); err != nil {
						errorf("udp", "%v", err)
					} else if N, err := udp.WriteToUDP(msg, addr); err != nil {
						errorf("udp", "%v", err)
					} else {
						infof("udp", "sent %v bytes to %v", N, addr)
						if debug {
							infof("udp", "packet\n%s", codec.Dump(msg[0:N], " ...          "))
						}
					}
				}
			}

			ctx.DeviceList.Apply(f)
		}
	}

	for {
		request := make([]byte, 2048)

		if N, raddr, err := udp.ReadFromUDP(request); err != nil {
			return err
		} else {
			if debug {
				debugf("udp", "received %v bytes from %v\n%s", N, raddr, codec.Dump(request[0:N], " ...          "))
			}

			handle(raddr, request[:N])
		}
	}
}

func tcpListenAndServe(ctx *simulator.Context, c *net.TCPListener) error {
	handle := func(connection net.Conn) {
		addr := connection.RemoteAddr()
		packet := make([]byte, 2048)
		deadline := time.Now().Add(5 * time.Second)

		connection.SetDeadline(deadline)

		if N, err := connection.Read(packet); err != nil {
			errorf("tcp", "%v", err)
		} else {
			if debug {
				debugf("tcp", "received %v bytes from %v\n%s", N, addr, codec.Dump(packet[0:N], " ...          "))
			}

			if request, err := messages.UnmarshalRequest(packet[0:N]); err != nil {
				errorf("tcp", "%v", err)
			} else {
				f := func(s simulator.Simulator) {
					if response, err := s.Handle(request); err != nil {
						warnf(tag(request), "%v", err)
					} else if !isNil(response) {
						if msg, err := codec.Marshal(response); err != nil {
							errorf("tcp", "%v", err)
						} else if N, err := connection.Write(msg); err != nil {
							errorf("tcp", "%v", err)
						} else {
							infof("tcp", "sent %v bytes to %v", N, addr)
							if debug {
								infof("tcp", "packet\n%s", codec.Dump(msg[0:N], " ...          "))
							}
						}
					}
				}

				ctx.DeviceList.Apply(f)
			}
		}

		connection.Close()
	}

	for {
		if client, err := c.Accept(); err != nil {
			return err
		} else {
			handle(client)
		}
	}
}

func tick(ctx *simulator.Context) {
	f := func(s simulator.Simulator) {
		s.Tick()
	}

	ctx.DeviceList.Apply(f)
}

func tasks(ctx *simulator.Context) {
	f := func(s simulator.Simulator) {
		s.RunTasks()
	}

	ctx.DeviceList.Apply(f)
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
			errorf("udp", "failed to write to UDP socket [%v]", err)
		} else if debug {
			infof("udp", "sent %v bytes to %v\n%s", N, dest, codec.Dump(msg[0:N], " ...          "))
		}
	}
}

func debugf(tag any, format string, args ...any) {
	f := fmt.Sprintf("%-10v  %v", tag, format)

	log.Debugf(f, args...)
}

func infof(tag any, format string, args ...any) {
	f := fmt.Sprintf("%-10v  %v", tag, format)

	log.Infof(f, args...)
}

func warnf(tag any, format string, args ...any) {
	f := fmt.Sprintf("%-10v  %v", tag, format)

	log.Warnf(f, args...)
}

func errorf(tag any, format string, args ...any) {
	f := fmt.Sprintf("%-10v  %v", tag, format)

	log.Errorf(f, args...)
}

func isNil(v any) bool {
	if v == nil {
		return true
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Ptr,
		reflect.Map,
		reflect.Array,
		reflect.Chan,
		reflect.Slice:
		return reflect.ValueOf(v).IsNil()
	}

	return false
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
