package UT0311L04

import (
	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) getDoorControlState(request *messages.GetDoorControlStateRequest) (*messages.GetDoorControlStateResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	if request.Door < 1 || request.Door > 4 {
		return nil, nil
	}

	response := messages.GetDoorControlStateResponse{
		SerialNumber: s.SerialNumber,
		Door:         request.Door,
		ControlState: s.Doors.ControlState(request.Door),
		Delay:        s.Doors.Delay(request.Door).Seconds(),
	}

	return &response, nil
}
