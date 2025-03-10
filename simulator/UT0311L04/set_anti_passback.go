package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) setAntiPassback(request *messages.SetAntiPassbackRequest) (*messages.SetAntiPassbackResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	ok := false

	if request.AntiPassback <= 0x04 {
		s.AntiPassback = types.AntiPassback(request.AntiPassback)
		ok = true
	}

	response := messages.SetAntiPassbackResponse{
		SerialNumber: s.SerialNumber,
		Ok:           ok,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
