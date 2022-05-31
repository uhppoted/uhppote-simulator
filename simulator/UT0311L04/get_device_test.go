package UT0311L04

import (
	"net"
	"reflect"
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func TestGetDeviceWithMatchingAddress(t *testing.T) {
	MAC, _ := net.ParseMAC("00:66:19:39:55:2d")
	released, _ := types.DateFromString("2020-12-05")
	listener := net.UDPAddr{IP: net.IPv4(10, 0, 0, 10), Port: 43210}
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		IpAddress:    net.IPv4(10, 0, 0, 100),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MacAddress:   types.MacAddress(MAC),
		Version:      9876,
		Released:     (*ReleaseDate)(&released),
		Listener:     &listener,
		Cards:        entities.CardList{},
		Events:       entities.EventList{},
		Doors:        map[uint8]*entities.Door{},

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}
	date, _ := types.DateFromString("2020-12-05")

	expected := entities.Message{
		Destination: &src,
		Message: &messages.GetDeviceResponse{
			SerialNumber: 12345,
			IpAddress:    net.IPv4(10, 0, 0, 100),
			SubnetMask:   net.IPv4(255, 255, 255, 0),
			Gateway:      net.IPv4(10, 0, 0, 1),
			MacAddress:   types.MacAddress(MAC),
			Version:      9876,
			Date:         date,
		},
	}

	request := messages.GetDeviceRequest{
		SerialNumber: 12345,
	}

	s.getDevice(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("'get-device' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}
}

func TestGetDeviceWithAddress0(t *testing.T) {
	MAC, _ := net.ParseMAC("00:66:19:39:55:2d")
	released, _ := types.DateFromString("2020-12-05")
	listener := net.UDPAddr{IP: net.IPv4(10, 0, 0, 10), Port: 43210}
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		IpAddress:    net.IPv4(10, 0, 0, 100),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MacAddress:   types.MacAddress(MAC),
		Version:      9876,
		Released:     (*ReleaseDate)(&released),
		Listener:     &listener,
		Cards:        entities.CardList{},
		Events:       entities.EventList{},
		Doors:        map[uint8]*entities.Door{},

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}
	date, _ := types.DateFromString("2020-12-05")

	expected := entities.Message{
		Destination: &src,
		Message: &messages.GetDeviceResponse{
			SerialNumber: 12345,
			IpAddress:    net.IPv4(10, 0, 0, 100),
			SubnetMask:   net.IPv4(255, 255, 255, 0),
			Gateway:      net.IPv4(10, 0, 0, 1),
			MacAddress:   types.MacAddress(MAC),
			Version:      9876,
			Date:         date,
		},
	}

	request := messages.GetDeviceRequest{
		SerialNumber: 0,
	}

	s.getDevice(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("'get-device' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}
}

func TestGetDeviceWithDifferentAddress(t *testing.T) {
	MAC, _ := net.ParseMAC("00:66:19:39:55:2d")
	released, _ := types.DateFromString("2020-12-05")
	listener := net.UDPAddr{IP: net.IPv4(10, 0, 0, 10), Port: 43210}
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		IpAddress:    net.IPv4(10, 0, 0, 100),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MacAddress:   types.MacAddress(MAC),
		Version:      9876,
		Released:     (*ReleaseDate)(&released),
		Listener:     &listener,
		Cards:        entities.CardList{},
		Events:       entities.EventList{},
		Doors:        map[uint8]*entities.Door{},

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	request := messages.GetDeviceRequest{
		SerialNumber: 54321,
	}

	s.getDevice(&src, &request)

	select {
	case <-txq:
		t.Fatalf("'get-device' sent response to request with different request")
	default:
	}
}
