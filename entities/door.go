package entities

import (
	"encoding/json"
	"sync"
	"time"
)

type Delay time.Duration

type Door struct {
	ControlState    uint8 `json:"control"`
	Delay           Delay `json:"delay"`
	overrideState   uint8
	profileDisabled bool
	buttonDisabled  bool
	open            bool
	button          bool
	openTimer       *time.Timer
	unlockedUntil   *time.Time
	pressedUntil    *time.Time
	guard           sync.RWMutex
}

const (
	NormallyOpen   = uint8(1)
	NormallyClosed = uint8(2)
	Controlled     = uint8(3)
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

func NewDoor(id uint8) *Door {
	door := new(Door)

	door.ControlState = 3
	door.Delay = Delay(5 * 1000000000)
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

func (d *Door) Unlock(duration time.Duration) bool {
	if d == nil {
		return false
	}

	d.guard.Lock()
	defer d.guard.Unlock()

	if d.ControlState == NormallyClosed || d.overrideState == NormallyClosed {
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

func (d *Door) PressButton(duration time.Duration) (pressed bool, reason uint8) {
	pressed = false
	reason = 0x14

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
		reason = 0x1e
		return
	}

	if d.ControlState == NormallyClosed || d.overrideState == NormallyClosed {
		reason = 0x14
		return
	}

	unlockUntil := time.Now().UTC()
	unlockUntil = unlockUntil.Add(duration)
	unlockUntil = unlockUntil.Add(time.Duration(d.Delay))

	if d.unlockedUntil == nil || d.unlockedUntil.Before(unlockUntil) {
		d.unlockedUntil = &unlockUntil
	}

	reason = 0x00

	return
}

func (d *Door) OverrideState(state uint8) bool {
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

	return d.ControlState == NormallyClosed || d.overrideState == NormallyClosed
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

func (d *Door) unlocked() bool {
	switch d.overrideState {
	case NormallyOpen:
		return true

	case NormallyClosed:
		return false

	case Controlled:
		return d.unlockedUntil != nil && d.unlockedUntil.After(time.Now())

	default:
		switch d.ControlState {
		case NormallyOpen:
			return true

		case NormallyClosed:
			return false

		case Controlled:
			return d.unlockedUntil != nil && d.unlockedUntil.After(time.Now())
		}
	}

	return false
}

func (d *Door) pressed() bool {
	return d.pressedUntil != nil && d.pressedUntil.After(time.Now())
}
