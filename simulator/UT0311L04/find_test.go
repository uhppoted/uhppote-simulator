package UT0311L04

import (
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func TestFindDeviceWithMatchingAddress(t *testing.T) {
	MAC, _ := net.ParseMAC("00:66:19:39:55:2d")
	listener := net.UDPAddr{IP: net.IPv4(10, 0, 0, 10), Port: 43210}
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		IpAddress:    net.IPv4(10, 0, 0, 100),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MacAddress:   types.MacAddress(MAC),
		Version:      9876,
		Listener:     &listener,
		Cards:        entities.CardList{},
		Events:       entities.EventList{},
		Doors:        map[uint8]*entities.Door{},

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}
	utc := time.Now().UTC()
	datetime := utc.Add(time.Duration(s.TimeOffset))

	expected := entities.Message{
		Destination: &src,
		Message: &messages.FindDevicesResponse{
			SerialNumber: 12345,
			IpAddress:    net.IPv4(10, 0, 0, 100),
			SubnetMask:   net.IPv4(255, 255, 255, 0),
			Gateway:      net.IPv4(10, 0, 0, 1),
			MacAddress:   types.MacAddress(MAC),
			Version:      9876,
			Date:         types.Date(datetime),
		},
	}

	request := messages.FindDevicesRequest{
		SerialNumber: 12345,
	}

	s.find(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		if !reflect.DeepEqual(response.Destination, expected.Destination) {
			t.Errorf("'find' sent incorrect rsponse with incorrect destination address\n   expected: %+v\n   got:      %+v\n", expected.Destination, response.Destination)
		}

		// INTERIM HACK to compare messages with dates that differ in the sub-second range
		p := fmt.Sprintf("%v", response.Message)
		q := fmt.Sprintf("%v", expected.Message)
		if p != q {
			t.Errorf("'find' sent incorrect rsponse with incorrect message\n   expected: %+v\n   got:      %+v\n", expected.Message, response.Message)
		}
	}
}

func TestFindDeviceWithAddress0(t *testing.T) {
	MAC, _ := net.ParseMAC("00:66:19:39:55:2d")
	listener := net.UDPAddr{IP: net.IPv4(10, 0, 0, 10), Port: 43210}
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		IpAddress:    net.IPv4(10, 0, 0, 100),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MacAddress:   types.MacAddress(MAC),
		Version:      9876,
		Listener:     &listener,
		Cards:        entities.CardList{},
		Events:       entities.EventList{},
		Doors:        map[uint8]*entities.Door{},

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}
	utc := time.Now().UTC()
	datetime := utc.Add(time.Duration(s.TimeOffset))

	expected := entities.Message{
		Destination: &src,
		Message: &messages.FindDevicesResponse{
			SerialNumber: 12345,
			IpAddress:    net.IPv4(10, 0, 0, 100),
			SubnetMask:   net.IPv4(255, 255, 255, 0),
			Gateway:      net.IPv4(10, 0, 0, 1),
			MacAddress:   types.MacAddress(MAC),
			Version:      9876,
			Date:         types.Date(datetime),
		},
	}

	request := messages.FindDevicesRequest{
		SerialNumber: 0,
	}

	s.find(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		if !reflect.DeepEqual(response.Destination, expected.Destination) {
			t.Errorf("'find' sent incorrect rsponse with incorrect destination address\n   expected: %+v\n   got:      %+v\n", expected.Destination, response.Destination)
		}

		// INTERIM HACK to compare messages with dates that differ in the sub-second range
		p := fmt.Sprintf("%v", response.Message)
		q := fmt.Sprintf("%v", expected.Message)
		if p != q {
			t.Errorf("'find' sent incorrect rsponse with incorrect message\n   expected: %+v\n   got:      %+v\n", expected.Message, response.Message)
		}
	}
}

func TestFindDeviceWithDifferentAddress(t *testing.T) {
	MAC, _ := net.ParseMAC("00:66:19:39:55:2d")
	listener := net.UDPAddr{IP: net.IPv4(10, 0, 0, 10), Port: 43210}
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		IpAddress:    net.IPv4(10, 0, 0, 100),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MacAddress:   types.MacAddress(MAC),
		Version:      9876,
		Listener:     &listener,
		Cards:        entities.CardList{},
		Events:       entities.EventList{},
		Doors:        map[uint8]*entities.Door{},

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	request := messages.FindDevicesRequest{
		SerialNumber: 54321,
	}

	s.find(&src, &request)

	timeout := time.After(500 * time.Millisecond)

	select {
	case <-timeout:
	case <-txq:
		t.Fatal("'find' sent response to request with different request")
	}
}
