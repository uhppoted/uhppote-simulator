package UT0311L04

import (
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func (s *UT0311L04) unlockDoor(request *messages.OpenDoorRequest) (*messages.OpenDoorResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	if request.Door < 1 || request.Door > 4 {
		return nil, nil
	}

	door := request.Door
	granted := true

	s.Doors.Unlock(door, 0*time.Second)

	datetime := types.DateTime(time.Now().UTC().Add(time.Duration(s.TimeOffset)))
	event := entities.Event{
		Type:      0x02,
		Granted:   granted,
		Door:      door,
		Direction: 0x01,
		Card:      3922570474,
		Timestamp: datetime,
		Reason:    0x2c,
	}

	s.add(event)

	response := messages.OpenDoorResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    granted,
	}

	return &response, nil
}
