package UT0311L04

import (
	"net"
	"reflect"
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func TestRecordSpecialEvents(t *testing.T) {
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber:        12345,
		RecordSpecialEvents: false,

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	expected := entities.Message{
		Destination: &src,
		Message: &messages.RecordSpecialEventsResponse{
			SerialNumber: 12345,
			Succeeded:    true,
		},
	}

	request := messages.RecordSpecialEventsRequest{
		SerialNumber: 12345,
		Enable:       true,
	}

	s.recordSpecialEvents(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("'record-special-events' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}

	if !s.RecordSpecialEvents {
		t.Errorf("'record-special-events' failed to update simulator 'RecordSpecialEvents' field\n   expected: %+v\n   got:      %+v\n", true, s.RecordSpecialEvents)
	}
}

func TestRecordSpecialEventsDisable(t *testing.T) {
	txq := make(chan entities.Message, 8)

	s := UT0311L04{
		SerialNumber:        12345,
		RecordSpecialEvents: true,

		txq: txq,
	}

	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	expected := entities.Message{
		Destination: &src,
		Message: &messages.RecordSpecialEventsResponse{
			SerialNumber: 12345,
			Succeeded:    true,
		},
	}

	request := messages.RecordSpecialEventsRequest{
		SerialNumber: 12345,
		Enable:       false,
	}

	s.recordSpecialEvents(&src, &request)

	response := <-txq

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("'record-special-events' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
	}

	if s.RecordSpecialEvents {
		t.Errorf("'record-special-events' failed to update simulator 'RecordSpecialEvents' field\n   expected: %+v\n   got:      %+v\n", false, s.RecordSpecialEvents)
	}
}
