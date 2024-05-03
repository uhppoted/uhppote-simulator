package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func (s *UT0311L04) setDoorControlState(request *messages.SetDoorControlStateRequest) (*messages.SetDoorControlStateResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	if request.Door < 1 || request.Door > 4 {
		fmt.Printf("ERROR: Invalid door' - expected: [1..4], received:%d", request.Door)
		return nil, nil
	}

	door := request.Door

	s.Doors.SetControlState(door, request.ControlState)
	s.Doors.SetDelay(door, entities.Delay(uint64(request.Delay)*1000000000))

	response := messages.SetDoorControlStateResponse{
		SerialNumber: s.SerialNumber,
		Door:         door,
		ControlState: s.Doors.ControlState(door),
		Delay:        s.Doors.Delay(door).Seconds(),
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
