package UT0311L04

import (
	"net"
	"net/netip"
	"reflect"
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
)

// TODO: deferred pending some way to compare Date field
// func TestHandleFindDevices(t *testing.T) {
//  MAC, _ := net.ParseMAC("00:66:19:39:55:2d")
//  now := types.Date(time.Now().UTC())
//
//  request := messages.FindDevicesRequest{}
//
//  response := messages.FindDevicesResponse{
//      SerialNumber: 12345,
//      IpAddress:    net.IPv4(10, 0, 0, 100),
//      SubnetMask:   net.IPv4(255, 255, 255, 0),
//      Gateway:      net.IPv4(10, 0, 0, 1),
//      MacAddress:   types.MacAddress(MAC),
//      Version:      9876,
//      Date:         now,
//  }
//
//  testHandle(&request, &response, t)
// }

func TestGetDeviceWithMatchingAddress(t *testing.T) {
	MAC, _ := net.ParseMAC("00:66:19:39:55:2d")
	listener := netip.MustParseAddrPort("10.0.0.10:43210")

	s := UT0311L04{
		SerialNumber: 12345,
		IpAddress:    net.IPv4(10, 0, 0, 100),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MacAddress:   types.MacAddress(MAC),
		Version:      9876,
		Released:     types.MustParseDate("2020-12-05"),
		Listener:     listener,
		Cards:        entities.CardList{},
		Events:       entities.EventList{},
		Doors:        entities.MakeDoors(),
	}

	expected := messages.GetDeviceResponse{
		SerialNumber: 12345,
		IpAddress:    net.IPv4(10, 0, 0, 100),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MacAddress:   types.MacAddress(MAC),
		Version:      9876,
		Date:         types.MustParseDate("2020-12-05"),
	}

	request := messages.GetDeviceRequest{
		SerialNumber: 12345,
	}

	if response, err := s.getDevice(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response == nil {
		t.Fatalf("Invalid response (%v)", response)
	} else if !reflect.DeepEqual(*response, expected) {
		t.Errorf("'get-device' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}
}

func TestGetDeviceWithAddress0(t *testing.T) {
	MAC, _ := net.ParseMAC("00:66:19:39:55:2d")
	listener := netip.MustParseAddrPort("10.0.0.10:43210")

	s := UT0311L04{
		SerialNumber: 12345,
		IpAddress:    net.IPv4(10, 0, 0, 100),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MacAddress:   types.MacAddress(MAC),
		Version:      9876,
		Released:     types.MustParseDate("2020-12-05"),
		Listener:     listener,
		Cards:        entities.CardList{},
		Events:       entities.EventList{},
		Doors:        entities.MakeDoors(),
	}

	expected := messages.GetDeviceResponse{
		SerialNumber: 12345,
		IpAddress:    net.IPv4(10, 0, 0, 100),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MacAddress:   types.MacAddress(MAC),
		Version:      9876,
		Date:         types.MustParseDate("2020-12-05"),
	}

	request := messages.GetDeviceRequest{
		SerialNumber: 0,
	}

	if response, err := s.getDevice(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response == nil {
		t.Fatalf("Invalid response (%v)", response)
	} else if !reflect.DeepEqual(*response, expected) {
		t.Errorf("'get-device' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}
}

func TestGetDeviceWithDifferentAddress(t *testing.T) {
	MAC, _ := net.ParseMAC("00:66:19:39:55:2d")
	listener := netip.MustParseAddrPort("10.0.0.10:43210")

	s := UT0311L04{
		SerialNumber: 12345,
		IpAddress:    net.IPv4(10, 0, 0, 100),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MacAddress:   types.MacAddress(MAC),
		Version:      9876,
		Released:     types.MustParseDate("2020-12-05"),
		Listener:     listener,
		Cards:        entities.CardList{},
		Events:       entities.EventList{},
		Doors:        entities.MakeDoors(),
	}

	request := messages.GetDeviceRequest{
		SerialNumber: 54321,
	}

	if response, err := s.getDevice(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response != nil {
		t.Fatalf("'get-device' sent response to request with different request")
	}
}
