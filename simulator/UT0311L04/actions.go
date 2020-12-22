package UT0311L04

import (
	"fmt"
	"time"

	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
)

const (
	swipePass      uint8 = 0x01
	noPrivilege    uint8 = 0x06
	normallyClosed uint8 = 0x0b
	noPass         uint8 = 0x12
)

func (s *UT0311L04) Swipe(cardNumber uint32, door uint8) (bool, error) {
	if door < 1 || door > 4 {
		return false, fmt.Errorf("%v: invalid door %d", s.DeviceID(), door)
	}

	swiped := func(eventType uint8, granted bool, reason uint8) {
		datetime := time.Now().UTC().Add(time.Duration(s.TimeOffset))
		event := entities.Event{
			Type:       eventType,
			Granted:    granted,
			Door:       door,
			Direction:  1,
			CardNumber: cardNumber,
			Timestamp:  types.DateTime(datetime),
			Reason:     reason,
		}

		s.add(&event)
	}

	for _, c := range s.Cards {
		if c == nil || c.CardNumber != cardNumber {
			continue
		}

		if !c.Doors[door] {
			swiped(0x01, false, noPrivilege)
			return false, nil
		}

		if d, ok := s.Doors[door]; ok {
			if d.ControlState == entities.NormallyClosed {
				swiped(0x01, false, normallyClosed)
				return false, nil
			}

			if d.Unlock() {
				swiped(0x02, true, swipePass)
				return true, nil
			}
		}

		break
	}

	// Denied
	swiped(0x01, false, noPass)

	return false, nil
}

func (s *UT0311L04) Open(door uint8, duration *time.Duration) (bool, error) {
	if door < 1 || door > 4 {
		return false, fmt.Errorf("%v: invalid door %d", s.DeviceID(), door)
	}

	onOpen := func(reason uint8) {
		if s.RecordSpecialEvents {
			datetime := time.Now().UTC().Add(time.Duration(s.TimeOffset))
			event := entities.Event{
				Type:       0x02,
				Granted:    true,
				Door:       door,
				Direction:  1,
				CardNumber: 8,
				Timestamp:  types.DateTime(datetime),
				Reason:     reason,
			}

			s.add(&event)
		}
	}

	onClose := func() {
		if s.RecordSpecialEvents {
			datetime := time.Now().UTC().Add(time.Duration(s.TimeOffset))
			event := entities.Event{
				Type:       0x02,
				Granted:    true,
				Door:       door,
				Direction:  1,
				CardNumber: 9,
				Timestamp:  types.DateTime(datetime),
				Reason:     0x18,
			}

			s.add(&event)
		}
	}

	opened := s.Doors[door].Open(duration, onOpen, onClose)

	return opened, nil
}

func (s *UT0311L04) Close(door uint8) (bool, error) {
	if door < 1 || door > 4 {
		return false, fmt.Errorf("%v: invalid door %d", s.DeviceID(), door)
	}

	onClose := func() {
		if s.RecordSpecialEvents {
			datetime := time.Now().UTC().Add(time.Duration(s.TimeOffset))
			event := entities.Event{
				Type:       0x02,
				Granted:    true,
				Door:       door,
				Direction:  1,
				CardNumber: 9,
				Timestamp:  types.DateTime(datetime),
				Reason:     0x18,
			}

			s.add(&event)
		}
	}

	closed := s.Doors[door].Close(onClose)

	return closed, nil
}
