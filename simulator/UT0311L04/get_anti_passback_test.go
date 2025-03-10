package UT0311L04

import (
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestHandleGetAntiPassback(t *testing.T) {
	request := messages.GetAntiPassbackRequest{
		SerialNumber: 12345,
	}

	response := messages.GetAntiPassbackResponse{
		SerialNumber: 12345,
		AntiPassback: 0x04,
	}

	testHandle(&request, &response, t)
}
