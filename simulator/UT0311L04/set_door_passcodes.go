package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) setDoorPasscodes(request *messages.SetDoorPasscodesRequest) (*messages.SetDoorPasscodesResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	ok := s.Doors.SetPasscodes(request.Door, request.Passcode1, request.Passcode2, request.Passcode3, request.Passcode4)

	response := messages.SetDoorPasscodesResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    ok,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
