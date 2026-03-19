package UT0311L04

import (
	"fmt"
	"time"

	"github.com/uhppoted/uhppote-core/messages"
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

	ok := s.Doors.SetFirstCard(request.Door, request.StartTime, request.EndTime, request.StartDoorControl, request.EndDoorControl, weekdays)

	response := messages.SetFirstCardResponse{
		SerialNumber: s.SerialNumber,
		Ok:           ok,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
