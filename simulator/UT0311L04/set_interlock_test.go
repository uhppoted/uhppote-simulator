package UT0311L04

import (
	"net"
	"reflect"
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func TestSetInterlock1(t *testing.T) {
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		Doors:        entities.MakeDoors(),

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	expected := entities.Message{
		Destination: &src,
		Message: &messages.SetInterlockResponse{
			SerialNumber: 12345,
			Succeeded:    true,
		},
	}

	request := messages.SetInterlockRequest{
		SerialNumber: 12345,
		Interlock:    1,
	}

	s.setInterlock(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("'set-interlock' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}

	if s.Doors.Interlock != 1 {
		t.Errorf("'set-interlock' failed to update simulator 'interlock' field\n   expected: %+v\n   got:      %+v\n", 1, s.Doors.Interlock)
	}
}

func TestSetInterlock2(t *testing.T) {
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		Doors:        entities.MakeDoors(),

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	expected := entities.Message{
		Destination: &src,
		Message: &messages.SetInterlockResponse{
			SerialNumber: 12345,
			Succeeded:    true,
		},
	}

	request := messages.SetInterlockRequest{
		SerialNumber: 12345,
		Interlock:    2,
	}

	s.setInterlock(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("'set-interlock' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}

	if s.Doors.Interlock != 2 {
		t.Errorf("'set-interlock' failed to update simulator 'interlock' field\n   expected: %+v\n   got:      %+v\n", 2, s.Doors.Interlock)
	}
}

func TestSetInterlock3(t *testing.T) {
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		Doors:        entities.MakeDoors(),

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	expected := entities.Message{
		Destination: &src,
		Message: &messages.SetInterlockResponse{
			SerialNumber: 12345,
			Succeeded:    true,
		},
	}

	request := messages.SetInterlockRequest{
		SerialNumber: 12345,
		Interlock:    3,
	}

	s.setInterlock(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("'set-interlock' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}

	if s.Doors.Interlock != 3 {
		t.Errorf("'set-interlock' failed to update simulator 'interlock' field\n   expected: %+v\n   got:      %+v\n", 3, s.Doors.Interlock)
	}
}

func TestSetInterlock4(t *testing.T) {
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		Doors:        entities.MakeDoors(),

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	expected := entities.Message{
		Destination: &src,
		Message: &messages.SetInterlockResponse{
			SerialNumber: 12345,
			Succeeded:    true,
		},
	}

	request := messages.SetInterlockRequest{
		SerialNumber: 12345,
		Interlock:    4,
	}

	s.setInterlock(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("'set-interlock' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}

	if s.Doors.Interlock != 4 {
		t.Errorf("'set-interlock' failed to update simulator 'interlock' field\n   expected: %+v\n   got:      %+v\n", 4, s.Doors.Interlock)
	}
}

func TestSetInterlockDisable(t *testing.T) {
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		Doors:        entities.MakeDoors(),

		txq: txq,
	}

	s.Doors.Interlock = 1

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	expected := entities.Message{
		Destination: &src,
		Message: &messages.SetInterlockResponse{
			SerialNumber: 12345,
			Succeeded:    true,
		},
	}

	request := messages.SetInterlockRequest{
		SerialNumber: 12345,
		Interlock:    0,
	}

	s.setInterlock(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("'set-interlock' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}

	if s.Doors.Interlock != 0 {
		t.Errorf("'set-interlock' failed to update simulator 'interlock' field\n   expected: %+v\n   got:      %+v\n", 0, s.Doors.Interlock)
	}
}

func TestSetInvalidInterlock(t *testing.T) {
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber: 12345,
		Doors:        entities.MakeDoors(),

		txq: txq,
	}

	s.Doors.Interlock = 3

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	expected := entities.Message{
		Destination: &src,
		Message: &messages.SetInterlockResponse{
			SerialNumber: 12345,
			Succeeded:    false,
		},
	}

	request := messages.SetInterlockRequest{
		SerialNumber: 12345,
		Interlock:    5,
	}

	s.setInterlock(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("'set-interlock' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}

	if s.Doors.Interlock != 3 {
		t.Errorf("'set-interlock' erroneously updated simulator 'interlock' field\n   expected: %+v\n   got:      %+v\n", 3, s.Doors.Interlock)
	}
}
