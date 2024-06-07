package UT0311L04

import (
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestHandleSetEventIndex(t *testing.T) {
	request := messages.SetEventIndexRequest{
		SerialNumber: 12345,
		Index:        7,
		MagicWord:    0x55aaaa55,
	}

	response := messages.SetEventIndexResponse{
		SerialNumber: 12345,
		Changed:      true,
	}

	testHandle(&request, &response, t)
}

func TestHandleGetEventIndex(t *testing.T) {
	request := messages.GetEventIndexRequest{
		SerialNumber: 12345,
	}

	response := messages.GetEventIndexResponse{
		SerialNumber: 12345,
		Index:        123,
	}

	testHandle(&request, &response, t)
}
