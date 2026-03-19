package UT0311L04

import (
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func TestHandleSetFirstCard(t *testing.T) {
	request := messages.SetFirstCardRequest{
		SerialNumber:     405419896,
		Door:             3,
		StartTime:        types.MustParseHHmm("08:30"),
		StartDoorControl: 1,
		EndTime:          types.MustParseHHmm("17:45"),
		EndDoorControl:   2,
		Monday:           true,
		Tuesday:          true,
		Wednesday:        false,
		Thursday:         true,
		Friday:           false,
		Saturday:         true,
		Sunday:           true,
	}

	response := messages.SetFirstCardResponse{
		SerialNumber: 405419896,
		Ok:           true,
	}

	testHandle(&request, &response, t)
}
