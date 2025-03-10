package UT0311L04

import (
	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) getAntiPassback(request *messages.GetAntiPassbackRequest) (*messages.GetAntiPassbackResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	response := messages.GetAntiPassbackResponse{
		SerialNumber: s.SerialNumber,
		AntiPassback: uint8(s.AntiPassback),
	}

	return &response, nil
}
