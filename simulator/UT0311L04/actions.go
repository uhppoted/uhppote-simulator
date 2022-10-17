package UT0311L04

import (
	"fmt"
	"time"

	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
)

const (
	swipePass       uint8 = 0x01
	noPrivilege     uint8 = 0x06
	normallyClosed  uint8 = 0x0b
	invalidTimezone uint8 = 0x0f
	noPass          uint8 = 0x12
)

// Implements the REST 'swipe' API.
//
// Checks the device and card permissions and unlocks the associated door
// if appropriate, but does not simulate opening the door. A 'swiped' event
// is generated and sent to the configured event listener (if any).
func (s *UT0311L04) Swipe(cardNumber uint32, door uint8) (bool, error) {
	if door < 1 || door > 4 {
		return false, fmt.Errorf("%v: invalid door %d", s.DeviceID(), door)
	}

	swiped := func(eventType uint8, granted bool, reason uint8) {
		datetime := types.DateTime(time.Now().UTC().Add(time.Duration(s.TimeOffset)))
		event := entities.Event{
			Type:      eventType,
			Granted:   granted,
			Door:      door,
			Direction: 1,
			Card:      cardNumber,
			Timestamp: &datetime,
			Reason:    reason,
		}

		s.add(event)
	}

	for _, c := range s.Cards {
		if c == nil || c.CardNumber != cardNumber {
			continue
		}

		profileID := c.Doors[door]

		// no access rights?
		if profileID < 1 || profileID > 254 {
			swiped(0x01, false, noPrivilege)
			return false, nil
		}

		// check against time profile
		if profileID >= 2 && profileID <= 254 {
			if d, ok := s.Doors[door]; ok {
				if d.IsProfileDisabled() {
					swiped(0x01, false, invalidTimezone)
					return false, nil
				}
			}

			if !s.checkTimeProfile(profileID) {
				swiped(0x01, false, invalidTimezone)
				return false, nil
			}
		}

		// unlock door
		if d, ok := s.Doors[door]; ok {
			if d.IsNormallyClosed() {
				swiped(0x01, false, normallyClosed)
				return false, nil
			}

			if d.Unlock(0 * time.Second) {
				swiped(0x02, true, swipePass)
				return true, nil
			}
		}

		break
	}

	// Denied!
	swiped(0x01, false, noPass)

	return false, nil
}

// Implements the REST 'open door' API.
//
// Checks the device and opens the door if has been unlocked by a simulated card
// swipe, associated button press or is configurated as normally open. The door
// is closed again after the open duration. A 'nil' duration will keep the door
// open until a 'door close' action closes it.
//
// A 'door open' event is generated and sent to the event listener (if any) if
// the door was opened and 'record special events' is enabled.
func (s *UT0311L04) Open(door uint8, duration *time.Duration) (bool, error) {
	if door < 1 || door > 4 {
		return false, fmt.Errorf("%v: invalid door %d", s.DeviceID(), door)
	}

	onOpen := func(reason uint8) {
		if s.RecordSpecialEvents {
			datetime := types.DateTime(time.Now().UTC().Add(time.Duration(s.TimeOffset)))
			event := entities.Event{
				Type:      0x02,
				Granted:   true,
				Door:      door,
				Direction: 1,
				Card:      8,
				Timestamp: &datetime,
				Reason:    reason,
			}

			s.add(event)
		}
	}

	onClose := func() {
		if s.RecordSpecialEvents {
			datetime := types.DateTime(time.Now().UTC().Add(time.Duration(s.TimeOffset)))
			event := entities.Event{
				Type:      0x02,
				Granted:   true,
				Door:      door,
				Direction: 1,
				Card:      9,
				Timestamp: &datetime,
				Reason:    0x18,
			}

			s.add(event)
		}
	}

	opened := s.Doors[door].Open(duration, onOpen, onClose)

	return opened, nil
}

// Implements the REST 'close door' API.
//
// Checks the device and closes the door if is has been opened by a simulated
// open door.
//
// A 'door close' event is generated and sent to the event listener (if any) if
// the door was closed and 'record special events' is enabled.
func (s *UT0311L04) Close(door uint8) (bool, error) {
	if door < 1 || door > 4 {
		return false, fmt.Errorf("%v: invalid door %d", s.DeviceID(), door)
	}

	onClose := func() {
		if s.RecordSpecialEvents {
			datetime := types.DateTime(time.Now().UTC().Add(time.Duration(s.TimeOffset)))
			event := entities.Event{
				Type:      0x02,
				Granted:   true,
				Door:      door,
				Direction: 1,
				Card:      9,
				Timestamp: &datetime,
				Reason:    0x18,
			}

			s.add(event)
		}
	}

	closed := s.Doors[door].Close(onClose)

	return closed, nil
}

// Implements the REST 'press button' API.
//
// Checks the device and unlocks the associated door as long as it is not configured
// as normally closed. The button is released the specified duration, with the door
// being held unlocked while the button is pressed.
//
// A 'button pressed' event is generated and sent to the event listener (if any) if
// 'record special events' is enabled. An event will not be generated if a previous
// button press is still active but will extend the duration of the action.
func (s *UT0311L04) ButtonPressed(door uint8, duration time.Duration) (bool, error) {
	if door < 1 || door > 4 {
		return false, fmt.Errorf("%v: invalid door %d", s.DeviceID(), door)
	}

	onUnlocked := func() {
		if s.RecordSpecialEvents {
			datetime := types.DateTime(time.Now().UTC().Add(time.Duration(s.TimeOffset)))
			event := entities.Event{
				Type:      0x02,
				Granted:   true,
				Door:      door,
				Direction: 1,
				Card:      1,
				Timestamp: &datetime,
				Reason:    0x14,
			}

			s.add(event)
		}
	}

	onNotUnlocked := func(reason uint8) {
		if s.RecordSpecialEvents {
			datetime := types.DateTime(time.Now().UTC().Add(time.Duration(s.TimeOffset)))
			event := entities.Event{
				Type:      0x03,
				Granted:   false,
				Door:      door,
				Direction: 1,
				Card:      6,
				Timestamp: &datetime,
				Reason:    reason,
			}

			s.add(event)
		}
	}

	pressed, reason := s.Doors[door].PressButton(duration)

	if pressed {
		if reason == 0x00 {
			onUnlocked()
		} else {
			onNotUnlocked(reason)
		}
	}

	return pressed && (reason == 0x00), nil
}

// Builds list of linked time profiles and then checks each individual profile
func (s *UT0311L04) checkTimeProfile(profileID uint8) bool {
	profiles := map[uint8]bool{}

	if profile, ok := s.TimeProfiles[profileID]; ok {
		profiles[profileID] = true

		linked := profile.LinkedProfileID
		for {
			if linked < 2 || linked > 254 || profiles[linked] {
				break
			}

			profiles[linked] = true
			if profile, ok := s.TimeProfiles[linked]; ok {
				linked = profile.LinkedProfileID
			}
		}

		for pid, _ := range profiles {
			if profile, ok := s.TimeProfiles[pid]; !ok {
				return true // IRL, a controller seems to default to ok if a linked time profile is not present
			} else if checkTimeProfile(profile, s.TimeOffset) {
				return true
			}
		}

		return false
	}

	return true // IRL, a controller seems to default to ok if time profile is not present
}

func checkTimeProfile(profile types.TimeProfile, offset entities.Offset) bool {
	utc := time.Now().UTC()
	adjusted := utc.Add(time.Duration(offset))
	now := types.HHmmFromTime(adjusted)
	today := types.Date(adjusted)
	weekday := today.Weekday()

	if profile.From == nil || profile.From.After(today) {
		return false
	}

	if profile.To == nil || profile.To.Before(today) {
		return false
	}

	if !profile.Weekdays[weekday] {
		return false
	}

	for _, i := range []uint8{1, 2, 3} {
		if segment, ok := profile.Segments[i]; ok {
			if !segment.Start.After(now) && !segment.End.Before(now) {
				return true
			}
		}
	}

	return false
}
