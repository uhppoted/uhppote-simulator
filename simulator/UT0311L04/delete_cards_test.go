package UT0311L04

import (
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestHandleDeleteCardsRequest(t *testing.T) {
	request := messages.DeleteCardsRequest{
		SerialNumber: 12345,
		MagicWord:    0x55aaaa55,
	}

	response := messages.DeleteCardsResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}
