package UT0311L04

import (
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestHandleGetCardsRequest(t *testing.T) {
	request := messages.GetCardsRequest{
		SerialNumber: 12345,
	}

	response := messages.GetCardsResponse{
		SerialNumber: 12345,
		Records:      3,
	}

	testHandle(&request, &response, t)
}
