package UT0311L04

import (
	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
	"net"
	"time"
)

func (s *UT0311L04) openDoor(addr *net.UDPAddr, request *messages.OpenDoorRequest) {
	if s.SerialNumber == request.SerialNumber {
		granted := false
		direction := uint8(0x01)
		door := request.Door

		if !(door < 1 || door > 4) {
			granted = true
			direction = s.Doors[door].Open()

			response := messages.OpenDoorResponse{
				SerialNumber: s.SerialNumber,
				Succeeded:    granted,
			}

			s.send(addr, &response)

			datetime := time.Now().UTC().Add(time.Duration(s.TimeOffset))
			event := entities.Event{
				Type:       0x02,
				Granted:    granted,
				Door:       door,
				Direction:  direction,
				CardNumber: 3922570474,
				Timestamp:  types.DateTime(datetime),
				Reason:     0x2c,
			}

			s.add(&event)
		}
	}
}
