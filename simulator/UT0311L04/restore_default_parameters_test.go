package UT0311L04

import (
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func TestRestoreDefaultParameters(t *testing.T) {
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 405419896,
		IpAddress:    net.IPv4(192, 168, 1, 100),
		SubnetMask:   net.IPv4(255, 254, 253, 252),
		Gateway:      net.IPv4(192, 168, 1, 1),
		Listener:     &net.UDPAddr{IP: net.IPv4(192, 168, 1, 100), Port: 60001},

		Doors: entities.MakeDoors(),

		txq: txq,
	}

	s.Doors.SetControlState(1, entities.NormallyOpen)
	s.Doors.SetControlState(2, entities.NormallyClosed)
	s.Doors.SetControlState(3, entities.NormallyOpen)
	s.Doors.SetControlState(4, entities.NormallyClosed)

	s.Doors.SetDelay(1, 15)
	s.Doors.SetDelay(2, 15)
	s.Doors.SetDelay(3, 15)
	s.Doors.SetDelay(4, 15)

	s.Doors.SetPasscodes(1, 12345, 54321, 0, 7531)
	s.Doors.SetPasscodes(2, 12345, 54321, 0, 7531)
	s.Doors.SetPasscodes(3, 12345, 54321, 0, 7531)
	s.Doors.SetPasscodes(4, 12345, 54321, 0, 7531)

	s.Events.Add(entities.Event{
		Index:     1,
		Type:      2,
		Granted:   true,
		Door:      3,
		Direction: 1,
		Card:      10058400,
		// 	// Timestamp : types.DateTimeNow(),
		// 	// Reason   : 15,
	})

	s.Events.SetIndex(1)

	s.Cards.Put(&entities.Card{
		CardNumber: 10058400,
		From:       types.ToDate(2024, time.January, 1),
		To:         types.ToDate(2024, time.December, 21),
		Doors: map[uint8]uint8{
			1: 1,
			2: 1,
			3: 0,
			4: 1,
		},
		PIN: 7531,
	})

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	expected := struct {
		response   entities.Message
		IpAddress  net.IP
		SubnetMask net.IP
		Gateway    net.IP
		Listener   *net.UDPAddr
	}{
		response: entities.Message{
			Destination: &src,
			Message: &messages.RestoreDefaultParametersResponse{
				SerialNumber: 405419896,
				Succeeded:    true,
			},
		},

		IpAddress:  net.IPv4(0, 0, 0, 0),
		SubnetMask: net.IPv4(255, 255, 255, 0),
		Gateway:    net.IPv4(0, 0, 0, 0),
		Listener:   nil,
	}

	request := messages.RestoreDefaultParametersRequest{
		SerialNumber: 405419896,
		MagicWord:    0x55aaaa55,
	}

	s.restoreDefaultParameters(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected.response) {
		t.Errorf("'restore-default-parameters' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected.response, response)
	}

	if !reflect.DeepEqual(s.IpAddress, expected.IpAddress) {
		t.Errorf("'restore-default-parameters' failed to update simulator IPv4 address\n   expected: %+v\n   got:      %+v\n", expected.IpAddress, s.IpAddress)
	}

	if !reflect.DeepEqual(s.Gateway, expected.Gateway) {
		t.Errorf("'restore-default-parameters' failed to update simulator IPv4 gateway\n   expected: %+v\n   got:      %+v\n", expected.Gateway, s.Gateway)
	}

	if s.Listener != nil {
		t.Errorf("'restore-default-parameters' failed to update simulator event listener\n   expected: %+v\n   got:      %+v\n", nil, s.Listener)
	}

	if s.Events.Size() != 0 {
		t.Errorf("'restore-default-parameters' failed to clear simulator event list\n   expected: %+v\n   got:      %+v\n", 0, s.Events.Size())
	}

	if s.Events.GetIndex() != 0 {
		t.Errorf("'restore-default-parameters' failed to reset simulator event index\n   expected: %+v\n   got:      %+v\n", 0, s.Events.GetIndex())
	}

	if s.Cards.Size() != 0 {
		t.Errorf("'restore-default-parameters' failed to clear simulator cards list\n   expected: %+v\n   got:      %+v\n", 0, s.Cards.Size())
	}

	for _, door := range []uint8{1, 2, 3, 4} {
		if s.Doors.ControlState(door) != entities.Controlled {
			t.Errorf("'restore-default-parameters' failed to reset door %v mode\n   expected: %+v\n   got:      %+v\n", door, entities.Controlled, s.Doors.ControlState(door))
		}

		if s.Doors.Delay(door) != 5 {
			t.Errorf("'restore-default-parameters' failed to reset door %v delay\n   expected: %+v\n   got:      %+v\n", door, 5, s.Doors.Delay(door))
		}

		if !reflect.DeepEqual(s.Doors.Passcodes(door), []uint32{0, 0, 0, 0}) {
			t.Errorf("'restore-default-parameters' failed to reset door %v passcodes\n   expected: %+v\n   got:      %+v\n", door, []uint32{0, 0, 0, 0}, s.Doors.Passcodes(door))
		}
	}

	if !reflect.DeepEqual(s.SubnetMask, expected.SubnetMask) {
		t.Errorf("'restore-default-parameters' failed to update simulator IPv4 netmask\n   expected: %+v\n   got:      %+v\n", expected.SubnetMask, s.SubnetMask)
	}
}
