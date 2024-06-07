package UT0311L04

import (
	"testing"
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func TestHandleGetEvent(t *testing.T) {
	datetime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-08-01 12:34:56", time.Local)
	timestamp := types.DateTime(datetime)

	request := messages.GetEventRequest{
		SerialNumber: 12345,
		Index:        2,
	}

	response := messages.GetEventResponse{
		SerialNumber: 12345,
		Index:        2,
		Type:         0x06,
		Granted:      true,
		Door:         4,
		Direction:    0x01,
		CardNumber:   555444321,
		Timestamp:    timestamp,
		Reason:       9,
	}

	testHandle(&request, &response, t)
}
