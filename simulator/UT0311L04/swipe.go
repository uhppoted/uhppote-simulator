package UT0311L04

import (
	"time"

	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func (s *UT0311L04) Swipe(deviceID uint32, cardNumber uint32, door uint8) (bool, uint32) {
	granted := false
	direction := uint8(0x01)
	eventType := uint8(0x01)
	reason := uint8(0x06)

	if s.SerialNumber == types.SerialNumber(deviceID) {
		for _, c := range s.Cards {
			if c != nil && c.CardNumber == cardNumber {
				if c.Doors[door] {
					granted = true
					direction = s.Doors[door].Open()
					eventType = 0x02
					reason = 0x01
				}
			}
		}
	}

	datetime := time.Now().UTC().Add(time.Duration(s.TimeOffset))
	event := entities.Event{
		Type:       eventType,
		Granted:    granted,
		Door:       door,
		Direction:  direction,
		CardNumber: cardNumber,
		Timestamp:  types.DateTime(datetime),
		Reason:     reason,
	}

	eventID := s.add(&event)

	return granted, eventID
}
