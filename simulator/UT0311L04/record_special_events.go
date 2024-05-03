package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) recordSpecialEvents(request *messages.RecordSpecialEventsRequest) (*messages.RecordSpecialEventsResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	s.RecordSpecialEvents = request.Enable

	response := messages.RecordSpecialEventsResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    true,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
