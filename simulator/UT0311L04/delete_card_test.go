package UT0311L04

import (
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestHandleDeleteCardRequest(t *testing.T) {
	request := messages.DeleteCardRequest{
		SerialNumber: 12345,
		CardNumber:   192837465,
	}

	response := messages.DeleteCardResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}
