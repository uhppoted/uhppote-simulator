package entities

import (
	"encoding/json"
	"slices"
	"sync"
	"time"

	"github.com/uhppoted/uhppote-core/types"
)

type Delay time.Duration
type Direction uint8

type Door struct {
	ControlState types.ControlState `json:"control"`
	Delay        Delay              `json:"delay"`
	Passcodes    []uint32           `json:"passcodes,omitempty"`
	FirstCard    *types.FirstCard   `json:"firstcard,omitempty"`

	overrideState   types.ControlState
	profileDisabled bool
	buttonDisabled  bool
	open            bool
	button          bool
	firstcardSwiped bool

	openTimer     *time.Timer
	unlockedUntil *time.Time
	pressedUntil  *time.Time
	pending       *types.FirstCard
	guard         sync.RWMutex
}

const (
	DirectionIn  Direction = 1
	DirectionOut Direction = 2
)

func (delay Delay) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(delay).String())
}

func (delay *Delay) UnmarshalJSON(bytes []byte) error {
	var s string

	err := json.Unmarshal(bytes, &s)
	if err != nil {
		return err
	}

	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}

	*delay = Delay(d)

	return nil
}

func (delay Delay) Seconds() uint8 {
	return uint8(time.Duration(delay).Seconds())
}

func DelayFromSeconds(t uint8) Delay {
	return Delay(time.Duration(int64(t) * 1000000000))
}

func (d Direction) String() string {
	switch d {
	case DirectionIn:
		return "in"
	case DirectionOut:
		return "out"
	default:
		return "(unknown)"
	}
}

func NewDoor(id uint8) *Door {
	door := new(Door)

	door.ControlState = 3
	door.Delay = Delay(5 * 1000000000)
	door.Passcodes = []uint32{0, 0, 0, 0}
	door.open = false
	door.button = false
	door.unlockedUntil = nil

	return door
}

func (d *Door) Open(duration *time.Duration, opened func(uint8), closed func()) bool {
	if d == nil {
		return false
	}

	d.guard.Lock()
	defer d.guard.Unlock()

	if !d.open && d.unlocked() {
		d.open = true

		go func() {
			if opened != nil {
				opened(0x17)
			}
		}()

		if duration != nil {
			d.openTimer = time.AfterFunc(*duration, func() {
				d.Close(closed)
			})
		}
	}

	return d.open
}

func (d *Door) Close(closed func()) bool {
	if d == nil {
		return false
	}

	d.guard.Lock()
	defer d.guard.Unlock()

	if d.open {
		d.open = false

		if d.openTimer != nil {
			d.openTimer.Stop()
		}

		go func() {
			if closed != nil {
				closed()
			}
		}()
	}

	return !d.open
}

func (d *Door) Unlock(duration time.Duration, firstcard bool) bool {
	if d == nil {
		return false
	}

	d.guard.Lock()
	defer d.guard.Unlock()

	if d.FirstCard != nil && firstcard {
		if !d.firstcardSwiped {
			d.firstcardSwiped = true
			d.ControlState = d.FirstCard.Active
		}
	}

	if d.ControlState == types.ModeNormallyClosed || d.overrideState == types.ModeNormallyClosed {
		return false
	}

	until := time.Now().UTC()
	until = until.Add(duration)
	until = until.Add(time.Duration(d.Delay))

	if d.unlockedUntil == nil || d.unlockedUntil.Before(until) {
		d.unlockedUntil = &until
	}

	return true
}

func (d *Door) UnlockWithPasscode(passcode uint32, duration time.Duration) bool {
	if d == nil {
		return false
	}

	if slices.Contains(d.Passcodes, passcode) {
		d.guard.Lock()
		defer d.guard.Unlock()

		// NTS: apparently not!
		// if d.ControlState == NormallyClosed || d.overrideState == NormallyClosed {
		// 	return false
		// }

		until := time.Now().UTC()
		until = until.Add(duration)
		until = until.Add(time.Duration(d.Delay))

		if d.unlockedUntil == nil || d.unlockedUntil.Before(until) {
			d.unlockedUntil = &until
		}

		return true
	}

	return false
}

func (d *Door) PressButton(duration time.Duration) (pressed bool, reason Reason) {
	pressed = false
	reason = ReasonPushbuttonOk

	if d == nil {
		return
	}

	d.guard.Lock()
	defer d.guard.Unlock()

	now := time.Now().UTC()
	pressUntil := time.Now().UTC()
	pressUntil = pressUntil.Add(time.Duration(d.Delay))

	if d.pressedUntil == nil || d.pressedUntil.Before(now) {
		pressed = true
	}

	if d.pressedUntil == nil || d.pressedUntil.Before(pressUntil) {
		d.pressedUntil = &pressUntil
	}

	if d.buttonDisabled {
		reason = ReasonPushbuttonDisabled
		return
	}

	if d.ControlState == types.ModeNormallyClosed || d.overrideState == types.ModeNormallyClosed {
		reason = ReasonPushbuttonOk
		return
	}

	unlockUntil := time.Now().UTC()
	unlockUntil = unlockUntil.Add(duration)
	unlockUntil = unlockUntil.Add(time.Duration(d.Delay))

	if d.unlockedUntil == nil || d.unlockedUntil.Before(unlockUntil) {
		d.unlockedUntil = &unlockUntil
	}

	reason = ReasonPushbuttonOk

	return
}

func (d *Door) OverrideState(state types.ControlState) bool {
	if d == nil {
		return false
	}

	d.guard.Lock()
	defer d.guard.Unlock()

	d.overrideState = state

	return true
}

func (d *Door) EnableProfile(enabled bool) bool {
	if d == nil {
		return false
	}

	d.guard.Lock()
	defer d.guard.Unlock()

	d.profileDisabled = !enabled

	return true
}

func (d *Door) EnableButton(enabled bool) bool {
	if d == nil {
		return false
	}

	d.guard.Lock()
	defer d.guard.Unlock()

	d.buttonDisabled = !enabled

	return true
}

func (d *Door) IsNormallyClosed() bool {
	if d == nil {
		return true
	}

	d.guard.RLock()
	defer d.guard.RUnlock()

	return d.ControlState == types.ModeNormallyClosed || d.overrideState == types.ModeNormallyClosed
}

func (d *Door) IsProfileDisabled() bool {
	if d == nil {
		return true
	}

	d.guard.RLock()
	defer d.guard.RUnlock()

	return d.profileDisabled
}

func (d *Door) IsOpen() bool {
	if d == nil {
		return false
	}

	d.guard.RLock()
	defer d.guard.RUnlock()

	return d.open
}

func (d *Door) IsUnlocked() bool {
	if d == nil {
		return false
	}

	d.guard.RLock()
	defer d.guard.RUnlock()

	return d.unlocked()
}

func (d *Door) IsButtonPressed() bool {
	if d == nil {
		return false
	}

	d.guard.RLock()
	defer d.guard.RUnlock()

	return d.pressed()
}

// Requires a card with first card privileges iff:
// 1. The door has an active first card configuration
// 2. The door has not been unlocked by a card with first card privileges
// 3. The first card configuration INACTIVE state is ModeFirstCardOnly
//
// i.e. don't require a card with first card privileges if any of the following are true:
// - the door does not have an active first card configuration
// - the doors first card configuration INACTIVE state is not ModeFirstCardOnly
// - the door has been unlocked by a card with first card privileges
func (d *Door) RequiresFirstCard() bool {
	if d == nil {
		return false
	}

	d.guard.RLock()
	defer d.guard.RUnlock()

	if d.FirstCard == nil {
		return false
	}

	if d.FirstCard.Inactive != types.ModeFirstCardOnly {
		return false
	}

	if d.firstcardSwiped {
		return false
	}

	now := time.Now()
	hhmm := types.HHmmFromTime(now)
	weekday := now.Weekday()

	if d.FirstCard.Weekdays[weekday] {
		if d.FirstCard.StartTime.Before(hhmm) && d.FirstCard.EndTime.After(hhmm) {
			return true
		}
	}

	return false
}

func (d *Door) unlocked() bool {
	switch d.overrideState {
	case types.ModeNormallyOpen:
		return true

	case types.ModeNormallyClosed:
		return false

	case types.ModeControlled:
		return d.unlockedUntil != nil && d.unlockedUntil.After(time.Now())

	default:
		switch d.ControlState {
		case types.ModeNormallyOpen:
			return true

		case types.ModeNormallyClosed:
			return false

		case types.ModeControlled:
			return d.unlockedUntil != nil && d.unlockedUntil.After(time.Now())
		}
	}

	return false
}

func (d *Door) pressed() bool {
	return d.pressedUntil != nil && d.pressedUntil.After(time.Now())
}

func (d *Door) MarshalJSON() ([]byte, error) {
	mode := uint8(0)
	switch d.ControlState {
	case types.ModeNormallyOpen:
		mode = 1
	case types.ModeNormallyClosed:
		mode = 2
	case types.ModeControlled:
		mode = 3
	}

	serializable := struct {
		Mode      uint8            `json:"control"`
		Delay     Delay            `json:"delay"`
		Passcodes []uint32         `json:"passcodes,omitempty"`
		FirstCard *types.FirstCard `json:"firstcard,omitempty"`
	}{
		Mode:      mode,
		Delay:     d.Delay,
		Passcodes: d.Passcodes,
		FirstCard: d.FirstCard,
	}

	return json.Marshal(serializable)
}

func (d *Door) UnmarshalJSON(bytes []byte) error {
	serializable := struct {
		Mode      uint8            `json:"control"`
		Delay     Delay            `json:"delay"`
		Passcodes []uint32         `json:"passcodes,omitempty"`
		FirstCard *types.FirstCard `json:"firstcard,omitempty"`
	}{}

	if err := json.Unmarshal(bytes, &serializable); err != nil {
		return err
	}

	switch serializable.Mode {
	case 1:
		d.ControlState = types.ModeNormallyOpen
	case 2:
		d.ControlState = types.ModeNormallyClosed
	case 3:
		d.ControlState = types.ModeControlled
	}

	d.Delay = serializable.Delay
	d.Passcodes = serializable.Passcodes
	d.FirstCard = serializable.FirstCard

	return nil
}
