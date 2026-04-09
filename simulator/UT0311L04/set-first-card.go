package UT0311L04

import (
	"fmt"
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) setFirstCard(request *messages.SetFirstCardRequest) (*messages.SetFirstCardResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	weekdays := map[time.Weekday]bool{
		time.Monday:    request.Monday,
		time.Tuesday:   request.Tuesday,
		time.Wednesday: request.Wednesday,
		time.Thursday:  request.Thursday,
		time.Friday:    request.Friday,
		time.Saturday:  request.Saturday,
		time.Sunday:    request.Sunday,
	}

	var active types.ControlState = types.ModeUnknown
	var inactive types.ControlState = types.ModeUnknown

	switch request.StartDoorControl {
	case 0:
		active = types.ModeControlled
	case 1:
		active = types.ModeNormallyOpen
	case 2:
		active = types.ModeNormallyClosed
	}

	switch request.EndDoorControl {
	case 0:
		inactive = types.ModeControlled
	case 1:
		inactive = types.ModeNormallyOpen
	case 2:
		inactive = types.ModeNormallyClosed
	case 3:
		inactive = types.ModeFirstCardOnly
	}

	ok := s.Doors.SetFirstCard(request.Door, request.StartTime, request.EndTime, active, inactive, weekdays)

	response := messages.SetFirstCardResponse{
		SerialNumber: s.SerialNumber,
		Ok:           ok,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
