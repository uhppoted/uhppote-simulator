package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) setDoorPasscodes(request *messages.SetDoorPasscodesRequest) (*messages.SetDoorPasscodesResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	response := messages.SetDoorPasscodesResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    false,
	}

	door := request.Door

	if !(door < 1 || door > 4) {
		s.Doors.SetPasscodes(door, request.Passcode1, request.Passcode2, request.Passcode3, request.Passcode4)
		response.Succeeded = true
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
