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
	bind, err := net.ResolveUDPAddr("udp4", ctx.BindAddress)
	if err != nil {
		log.Errorf("failed to resolve UDP bind address [%v]", err)
		return
	}

	connection, err := net.ListenUDP("udp", bind)
	if err != nil {
		log.Errorf("failed to bind to UDP socket [%v]", err)
		return
	}

	defer connection.Close()

	log.Infof("bound to address %s", bind)

	wait := make(chan int, 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go run(ctx, connection, wait)

	<-interrupt
	connection.Close()
	<-wait

	os.Exit(1)
}

func run(ctx *simulator.Context, connection *net.UDPConn, wait chan int) {
	go func() {
		err := listenAndServe(ctx, connection)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		wait <- 0
	}()

	go func() {
		for {
			msg := ctx.DeviceList.GetMessage()
			send(connection, msg.Destination, msg.Message)
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

func listenAndServe(ctx *simulator.Context, c *net.UDPConn) error {
	for {
		request, remote, err := receive(c)
		if err != nil {
			return err
		}

		handle(ctx, c, remote, request)
	}
}

func handle(ctx *simulator.Context, c *net.UDPConn, src *net.UDPAddr, bytes []byte) {
	request, err := messages.UnmarshalRequest(bytes)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	f := func(s simulator.Simulator) {
		s.Handle(src, request)
	}

	ctx.DeviceList.Apply(f)
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

func send(c *net.UDPConn, dest *net.UDPAddr, message interface{}) {
	msg, err := codec.Marshal(message)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	N, err := c.WriteTo(msg, dest)
	if err != nil {
		log.Errorf("failed to write to UDP socket [%v]", err)
	} else if debug {
		log.Infof("sent %v bytes to %v\n%s", N, dest, dump(msg[0:N], " ...          "))
	}
}

func dump(m []byte, prefix string) string {
	return regexp.MustCompile("(?m)^(.*)").ReplaceAllString(hex.Dump(m), prefix+"$1")
}
