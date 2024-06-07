package UT0311L04

import (
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func TestHandlePutCardRequest(t *testing.T) {
	request := messages.PutCardRequest{
		SerialNumber: 12345,
		CardNumber:   192837465,
		From:         types.MustParseDate("2019-01-01"),
		To:           types.MustParseDate("2019-12-31"),
		Door1:        1,
		Door2:        0,
		Door3:        1,
		Door4:        0,
	}

	response := messages.PutCardResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}
