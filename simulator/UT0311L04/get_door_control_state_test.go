package UT0311L04

import (
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestHandleGetDoorControlState(t *testing.T) {
	request := messages.GetDoorControlStateRequest{
		SerialNumber: 12345,
		Door:         2,
	}

	response := messages.GetDoorControlStateResponse{
		SerialNumber: 12345,
		Door:         2,
		ControlState: 2,
		Delay:        22,
	}

	testHandle(&request, &response, t)
}
