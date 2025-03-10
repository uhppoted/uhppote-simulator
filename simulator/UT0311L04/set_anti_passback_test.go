package UT0311L04

import (
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestHandleSetAntiPassback(t *testing.T) {
	request := messages.SetAntiPassbackRequest{
		SerialNumber: 12345,
		AntiPassback: 0x04,
	}

	response := messages.SetAntiPassbackResponse{
		SerialNumber: 12345,
		Ok:           true,
	}

	testHandle(&request, &response, t)
}
