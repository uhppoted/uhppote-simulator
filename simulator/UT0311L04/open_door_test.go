package UT0311L04

import (
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestHandleOpenDoor(t *testing.T) {
	request := messages.OpenDoorRequest{
		SerialNumber: 405419896,
		Door:         3,
	}

	response := messages.OpenDoorResponse{
		SerialNumber: 405419896,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}
