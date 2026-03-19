package UT0311L04

import (
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestHandleSetDoorControlState(t *testing.T) {
	request := messages.SetDoorControlStateRequest{
		SerialNumber: 405419896,
		Door:         2,
		ControlState: 3,
		Delay:        7,
	}

	response := messages.SetDoorControlStateResponse{
		SerialNumber: 405419896,
		Door:         2,
		ControlState: 3,
		Delay:        7,
	}

	testHandle(&request, &response, t)
}
