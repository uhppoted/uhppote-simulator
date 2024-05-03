package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) setInterlock(request *messages.SetInterlockRequest) (*messages.SetInterlockResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	ok := false

	if request.Interlock == 0 || request.Interlock == 1 || request.Interlock == 2 || request.Interlock == 3 || request.Interlock == 4 || request.Interlock == 8 {
		s.Doors.Interlock = request.Interlock
		ok = true
	}

	response := messages.SetInterlockResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    ok,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
