package UT0311L04

import (
	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) getDoorControlState(request *messages.GetDoorControlStateRequest) (*messages.GetDoorControlStateResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	if request.Door < 1 || request.Door > 4 {
		return nil, nil
	}

	mode := uint8(0)
	switch s.Doors.ControlState(request.Door) {
	case types.ModeNormallyOpen:
		mode = 1
	case types.ModeNormallyClosed:
		mode = 2
	case types.ModeControlled:
		mode = 3
	}

	response := messages.GetDoorControlStateResponse{
		SerialNumber: s.SerialNumber,
		Door:         request.Door,
		ControlState: mode,
		Delay:        s.Doors.Delay(request.Door).Seconds(),
	}

	return &response, nil
}
