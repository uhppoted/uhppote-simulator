package UT0311L04

import (
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestHandleOpenDoor(t *testing.T) {
	request := messages.OpenDoorRequest{
		SerialNumber: 12345,
		Door:         3,
	}

	response := messages.OpenDoorResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}
