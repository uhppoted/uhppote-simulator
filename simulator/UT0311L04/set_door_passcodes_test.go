package UT0311L04

import (
	"net"
	"reflect"
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func TestSetDoorPasscodes(t *testing.T) {
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 405419896,
		Doors:        entities.MakeDoors(),

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	expected := struct {
		response  entities.Message
		passcodes []uint32
	}{
		response: entities.Message{
			Destination: &src,
			Message: &messages.SetDoorPasscodesResponse{
				SerialNumber: 405419896,
				Succeeded:    true,
			},
		},

		passcodes: []uint32{12345, 0, 999999, 54321},
	}

	request := messages.SetDoorPasscodesRequest{
		SerialNumber: 405419896,
		Door:         3,
		Passcode1:    12345,
		Passcode2:    0,
		Passcode3:    999999,
		Passcode4:    54321,
	}

	s.setDoorPasscodes(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected.response) {
		t.Errorf("'set-door-passcodes' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected.response, response)
	}

	if !reflect.DeepEqual(s.Doors.Passcodes(3), expected.passcodes) {
		t.Errorf("'set-door-passcodes' failed to update simulator\n   expected: %+v\n   got:      %+v\n", expected.passcodes, s.Doors.Passcodes(3))
	}
}
