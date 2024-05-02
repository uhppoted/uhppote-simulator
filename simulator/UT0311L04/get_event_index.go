package UT0311L04

import (
	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) getEventIndex(request *messages.GetEventIndexRequest) (*messages.GetEventIndexResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	response := messages.GetEventIndexResponse{
		SerialNumber: s.SerialNumber,
		Index:        s.Events.GetIndex(),
	}

	return &response, nil
}
