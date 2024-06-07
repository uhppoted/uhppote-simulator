package UT0311L04

import (
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func TestHandleGetCardById(t *testing.T) {
	request := messages.GetCardByIDRequest{
		SerialNumber: 12345,
		CardNumber:   192837465,
	}

	response := messages.GetCardByIDResponse{
		SerialNumber: 12345,
		CardNumber:   192837465,
		From:         types.MustParseDate("2019-01-01"),
		To:           types.MustParseDate("2019-12-31"),
		Door1:        1,
		Door2:        0,
		Door3:        0,
		Door4:        1,
	}

	testHandle(&request, &response, t)
}

func TestHandleGetCardByIndex(t *testing.T) {
	request := messages.GetCardByIndexRequest{
		SerialNumber: 12345,
		Index:        2,
	}

	response := messages.GetCardByIndexResponse{
		SerialNumber: 12345,
		CardNumber:   192837465,
		From:         types.MustParseDate("2019-01-01"),
		To:           types.MustParseDate("2019-12-31"),
		Door1:        1,
		Door2:        0,
		Door3:        0,
		Door4:        1,
	}

	testHandle(&request, &response, t)
}
