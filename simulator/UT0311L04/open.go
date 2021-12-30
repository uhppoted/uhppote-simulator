package UT0311L04

import (
	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
	"net"
	"time"
)

func (s *UT0311L04) unlockDoor(addr *net.UDPAddr, request *messages.OpenDoorRequest) {
	if s.SerialNumber == request.SerialNumber {
		granted := false
		door := request.Door

		if !(door < 1 || door > 4) {
			granted = true
			s.Doors[door].Unlock(0 * time.Second)

			response := messages.OpenDoorResponse{
				SerialNumber: s.SerialNumber,
				Succeeded:    granted,
			}

			s.send(addr, &response)

			datetime := types.DateTime(time.Now().UTC().Add(time.Duration(s.TimeOffset)))
			event := entities.Event{
				Type:      0x02,
				Granted:   granted,
				Door:      door,
				Direction: 0x01,
				Card:      3922570474,
				Timestamp: &datetime,
				Reason:    0x2c,
			}

			s.add(event)
		}
	}
}
