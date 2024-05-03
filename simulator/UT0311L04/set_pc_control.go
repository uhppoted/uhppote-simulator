package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) setPCControl(request *messages.SetPCControlRequest) (*messages.SetPCControlResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	ok := false

	if request.MagicWord == 0x55aaaa55 {
		s.PCControl = request.Enable
		ok = true
	}

	response := messages.SetPCControlResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    ok,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
