package UT0311L04

import (
	"reflect"
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func TestActivateAccessKeypads(t *testing.T) {
	s := UT0311L04{
		SerialNumber: 405419896,
		Keypads: entities.Keypads{
			1: entities.KeypadNone,
			2: entities.KeypadNone,
			3: entities.KeypadNone,
			4: entities.KeypadNone,
		},
	}

	expected := &messages.ActivateAccessKeypadsResponse{
		SerialNumber: 405419896,
		Succeeded:    true,
	}

	request := messages.ActivateAccessKeypadsRequest{
		SerialNumber: 405419896,
		Reader1:      true,
		Reader2:      true,
		Reader3:      false,
		Reader4:      true,
	}

	if response, err := s.activateKeypads(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response == nil {
		t.Errorf("Incorrect activate-keypads response (%v)", response)
	} else {
		if !reflect.DeepEqual(response, expected) {
			t.Errorf("'activate-keypads' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected, response)
		}

		if s.Keypads[1] != entities.KeypadBoth && s.Keypads[2] != entities.KeypadBoth && s.Keypads[3] != entities.KeypadNone && s.Keypads[4] != entities.KeypadBoth {
			t.Errorf("'activate-keypads' failed to update simulator keypads fields\n   expected: %+v\n   got:      %+v\n",
				[]entities.Keypad{entities.KeypadBoth, entities.KeypadBoth, entities.KeypadNone, entities.KeypadBoth},
				[]entities.Keypad{s.Keypads[1], s.Keypads[2], s.Keypads[3], s.Keypads[4]})
		}
	}
}
