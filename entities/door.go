package entities

import (
	"encoding/json"
	"sync"
	"time"
)

type Delay time.Duration

type Door struct {
	ControlState  uint8 `json:"control"`
	Delay         Delay `json:"delay"`
	open          bool
	button        bool
	openTimer     *time.Timer
	unlockedUntil *time.Time
	guard         sync.RWMutex
}

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

	now := time.Now()

	d.guard.Lock()
	defer d.guard.Unlock()

	if !d.open && (d.ControlState == 0x01 || (d.ControlState == 0x03 && d.unlockedUntil != nil && d.unlockedUntil.After(now))) {
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
	if d != nil {
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

	return false
}

func (d *Door) Unlock() uint8 {
	if d != nil {
		now := time.Now().UTC()
		lockAt := now.Add(time.Duration(d.Delay))

		d.unlockedUntil = &lockAt

		return 0x01
	}

	return 0x00
}

func (d *Door) IsOpen() bool {
	if d != nil {
		d.guard.RLock()
		defer d.guard.RUnlock()

		return d.open
	}

	return false
}

func (d *Door) IsButtonPressed() bool {
	if d != nil {
		return false
	}

	return false
}
