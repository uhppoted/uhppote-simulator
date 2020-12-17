package UT0311L04

import (
	"fmt"
	"time"

	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func (s *UT0311L04) Swipe(deviceID uint32, cardNumber uint32, door uint8) (bool, uint32) {
	granted := false
	direction := uint8(0x01)
	eventType := uint8(0x01)
	reason := uint8(0x06)

	for _, c := range s.Cards {
		if c != nil && c.CardNumber == cardNumber {
			if c.Doors[door] {
				granted = true
				direction = s.Doors[door].Unlock()
				eventType = 0x02
				reason = 0x01
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

func (s *UT0311L04) Open(deviceID uint32, door uint8) (uint32, error) {
	if door < 1 || door > 4 {
		return 0, fmt.Errorf("%v: invalid doori %d", deviceID, door)
	}

	s.Doors[door].Open()

	var eventID uint32 = 0

	if s.RecordSpecialEvents {
		datetime := time.Now().UTC().Add(time.Duration(s.TimeOffset))
		event := entities.Event{
			Type:      0x02,
			Granted:   true,
			Door:      door,
			Direction: 1,
			Timestamp: types.DateTime(datetime),
			Reason:    0x17,
		}

		eventID = s.add(&event)
	}

	return eventID, nil
}

func (s *UT0311L04) Close(deviceID uint32, door uint8) (uint32, error) {
	if door < 1 || door > 4 {
		return 0, fmt.Errorf("%v: invalid doori %d", deviceID, door)
	}

	s.Doors[door].Close()

	var eventID uint32 = 0

	if s.RecordSpecialEvents {
		datetime := time.Now().UTC().Add(time.Duration(s.TimeOffset))
		event := entities.Event{
			Type:      0x02,
			Granted:   true,
			Door:      door,
			Direction: 1,
			Timestamp: types.DateTime(datetime),
			Reason:    0x18,
		}

		eventID = s.add(&event)
	}

	return eventID, nil
}
