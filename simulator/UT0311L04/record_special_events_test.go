package UT0311L04

import (
	"reflect"
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestRecordSpecialEvents(t *testing.T) {
	s := UT0311L04{
		SerialNumber:        12345,
		RecordSpecialEvents: false,
	}

	expected := messages.RecordSpecialEventsResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	request := messages.RecordSpecialEventsRequest{
		SerialNumber: 12345,
		Enable:       true,
	}

	if response, err := s.recordSpecialEvents(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response == nil {
		t.Fatalf("invalid response %v", response)
	} else {
		if !reflect.DeepEqual(*response, expected) {
			t.Errorf("'record-special-events' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, *response)
		}

		if !s.RecordSpecialEvents {
			t.Errorf("'record-special-events' failed to update simulator 'RecordSpecialEvents' field\n   expected: %+v\n   got:      %+v\n", true, s.RecordSpecialEvents)
		}
	}
}

func TestRecordSpecialEventsDisable(t *testing.T) {
	s := UT0311L04{
		SerialNumber:        12345,
		RecordSpecialEvents: true,
	}

	expected := messages.RecordSpecialEventsResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	request := messages.RecordSpecialEventsRequest{
		SerialNumber: 12345,
		Enable:       false,
	}

	if response, err := s.recordSpecialEvents(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response == nil {
		t.Fatalf("invalid response (%v)", response)
	} else {

		if !reflect.DeepEqual(*response, expected) {
			t.Errorf("'record-special-events' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, *response)
		}

		if s.RecordSpecialEvents {
			t.Errorf("'record-special-events' failed to update simulator 'RecordSpecialEvents' field\n   expected: %+v\n   got:      %+v\n", false, s.RecordSpecialEvents)
		}
	}
}
