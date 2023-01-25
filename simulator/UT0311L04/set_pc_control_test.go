package UT0311L04

import (
	"net"
	"reflect"
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func TestSetPCControl(t *testing.T) {
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		PCControl:    false,

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	expected := entities.Message{
		Destination: &src,
		Message: &messages.SetPCControlResponse{
			SerialNumber: 12345,
			Succeeded:    true,
		},
	}

	request := messages.SetPCControlRequest{
		SerialNumber: 12345,
		MagicWord:    0x55aaaa55,
		Enable:       true,
	}

	s.setPCControl(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("'set-pc-control' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}

	if !s.PCControl {
		t.Errorf("'set-pc-control' failed to update simulator 'PC control' field\n   expected: %+v\n   got:      %+v\n", true, s.PCControl)
	}
}

func TestSetPCControlDisable(t *testing.T) {
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		PCControl:    true,

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	expected := entities.Message{
		Destination: &src,
		Message: &messages.SetPCControlResponse{
			SerialNumber: 12345,
			Succeeded:    true,
		},
	}

	request := messages.SetPCControlRequest{
		SerialNumber: 12345,
		MagicWord:    0x55aaaa55,
		Enable:       false,
	}

	s.setPCControl(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("'set-pc-control' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}

	if s.PCControl {
		t.Errorf("'set-pc-control' failed to update simulator 'SetPCControl' field\n   expected: %+v\n   got:      %+v\n", false, s.PCControl)
	}
}

func TestSetPCControlWithoutMagicWord(t *testing.T) {
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		PCControl:    false,

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	expected := entities.Message{
		Destination: &src,
		Message: &messages.SetPCControlResponse{
			SerialNumber: 12345,
			Succeeded:    false,
		},
	}

	request := messages.SetPCControlRequest{
		SerialNumber: 12345,
		Enable:       true,
	}

	s.setPCControl(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("'set-pc-control' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}

	if s.PCControl {
		t.Errorf("'set-pc-control' incorrectly updated simulator 'PC control' field\n   expected: %+v\n   got:      %+v\n", false, s.PCControl)
	}
}
