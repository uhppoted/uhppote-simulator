package UT0311L04

import (
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) getTime(request *messages.GetTimeRequest) (*messages.GetTimeResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	utc := time.Now().UTC()
	datetime := utc.Add(time.Duration(s.TimeOffset))

	response := messages.GetTimeResponse{
		SerialNumber: s.SerialNumber,
		DateTime:     types.DateTime(datetime),
	}

	return &response, nil
}
