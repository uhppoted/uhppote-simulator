package UT0311L04

import (
	"reflect"
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func TestSetInterlock1(t *testing.T) {
	s := UT0311L04{
		SerialNumber: 12345,
		Doors:        entities.MakeDoors(),
	}

	expected := messages.SetInterlockResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	request := messages.SetInterlockRequest{
		SerialNumber: 12345,
		Interlock:    1,
	}

	if response, err := s.setInterlock(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response == nil {
		t.Fatalf("invalid response (%v)", response)
	} else {
		if !reflect.DeepEqual(*response, expected) {
			t.Errorf("'set-interlock' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, *response)
		}

		if s.Doors.Interlock != 1 {
			t.Errorf("'set-interlock' failed to update simulator 'interlock' field\n   expected: %+v\n   got:      %+v\n", 1, s.Doors.Interlock)
		}
	}
}

func TestSetInterlock2(t *testing.T) {
	s := UT0311L04{
		SerialNumber: 12345,
		Doors:        entities.MakeDoors(),
	}

	expected := messages.SetInterlockResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	request := messages.SetInterlockRequest{
		SerialNumber: 12345,
		Interlock:    2,
	}

	if response, err := s.setInterlock(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response == nil {
		t.Fatalf("invalid response (%v)", response)
	} else {
		if !reflect.DeepEqual(*response, expected) {
			t.Errorf("'set-interlock' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, *response)
		}

		if s.Doors.Interlock != 2 {
			t.Errorf("'set-interlock' failed to update simulator 'interlock' field\n   expected: %+v\n   got:      %+v\n", 2, s.Doors.Interlock)
		}
	}
}

func TestSetInterlock3(t *testing.T) {
	s := UT0311L04{
		SerialNumber: 12345,
		Doors:        entities.MakeDoors(),
	}

	expected := messages.SetInterlockResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	request := messages.SetInterlockRequest{
		SerialNumber: 12345,
		Interlock:    3,
	}

	if response, err := s.setInterlock(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response == nil {
		t.Fatalf("invalid response (%v)", response)
	} else {
		if !reflect.DeepEqual(*response, expected) {
			t.Errorf("'set-interlock' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
		}

		if s.Doors.Interlock != 3 {
			t.Errorf("'set-interlock' failed to update simulator 'interlock' field\n   expected: %+v\n   got:      %+v\n", 3, s.Doors.Interlock)
		}
	}
}

func TestSetInterlock4(t *testing.T) {
	s := UT0311L04{
		SerialNumber: 12345,
		Doors:        entities.MakeDoors(),
	}

	expected := messages.SetInterlockResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	request := messages.SetInterlockRequest{
		SerialNumber: 12345,
		Interlock:    4,
	}

	if response, err := s.setInterlock(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response == nil {
		t.Fatalf("invalid response (%v)", response)
	} else {
		if !reflect.DeepEqual(*response, expected) {
			t.Errorf("'set-interlock' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, *response)
		}

		if s.Doors.Interlock != 4 {
			t.Errorf("'set-interlock' failed to update simulator 'interlock' field\n   expected: %+v\n   got:      %+v\n", 4, s.Doors.Interlock)
		}
	}
}

func TestSetInterlockDisable(t *testing.T) {
	s := UT0311L04{
		SerialNumber: 12345,
		Doors:        entities.MakeDoors(),
	}

	s.Doors.Interlock = 1

	expected := messages.SetInterlockResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	request := messages.SetInterlockRequest{
		SerialNumber: 12345,
		Interlock:    0,
	}

	if response, err := s.setInterlock(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response == nil {
		t.Fatalf("invalid response (%v)", response)
	} else {
		if !reflect.DeepEqual(*response, expected) {
			t.Errorf("'set-interlock' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, *response)
		}

		if s.Doors.Interlock != 0 {
			t.Errorf("'set-interlock' failed to update simulator 'interlock' field\n   expected: %+v\n   got:      %+v\n", 0, s.Doors.Interlock)
		}
	}
}

func TestSetInvalidInterlock(t *testing.T) {
	s := UT0311L04{
		SerialNumber: 12345,
		Doors:        entities.MakeDoors(),
	}

	s.Doors.Interlock = 3

	expected := messages.SetInterlockResponse{
		SerialNumber: 12345,
		Succeeded:    false,
	}

	request := messages.SetInterlockRequest{
		SerialNumber: 12345,
		Interlock:    5,
	}

	if response, err := s.setInterlock(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response == nil {
		t.Fatalf("invalid response (%v)", response)
	} else {
		if !reflect.DeepEqual(*response, expected) {
			t.Errorf("'set-interlock' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, *response)
		}

		if s.Doors.Interlock != 3 {
			t.Errorf("'set-interlock' erroneously updated simulator 'interlock' field\n   expected: %+v\n   got:      %+v\n", 3, s.Doors.Interlock)
		}
	}
}
