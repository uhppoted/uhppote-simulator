package entities

import (
	"encoding/json"
	"time"
)

type Doors struct {
	Interlock uint8
	doors     map[uint8]*Door
}

func MakeDoors() Doors {
	return Doors{
		Interlock: 0,
		doors: map[uint8]*Door{
			1: NewDoor(1),
			2: NewDoor(2),
			3: NewDoor(3),
			4: NewDoor(4),
		},
	}
}

func (dd *Doors) SetControlState(door uint8, state uint8) {
	if d, ok := dd.doors[door]; ok {
		d.ControlState = state
	}
}

func (dd *Doors) SetDelay(door uint8, delay Delay) {
	if d, ok := dd.doors[door]; ok {
		d.Delay = delay
	}
}

func (dd *Doors) OverrideState(door uint8, state uint8) bool {
	if d, ok := dd.doors[door]; ok {
		return d.OverrideState(state)
	}

	return false
}

func (dd *Doors) EnableProfile(door uint8, enabled bool) bool {
	if d, ok := dd.doors[door]; ok {
		return d.EnableProfile(enabled)
	}

	return false
}

func (dd *Doors) Unlock(door uint8, duration time.Duration) bool {
	if d, ok := dd.doors[door]; ok {
		return d.Unlock(duration)
	}

	return false
}

func (dd *Doors) Open(door uint8, duration *time.Duration, opened func(uint8), closed func()) bool {
	if d, ok := dd.doors[door]; ok {
		return d.Open(duration, opened, closed)
	}

	return false
}

func (dd *Doors) Close(door uint8, closed func()) bool {
	if d, ok := dd.doors[door]; ok {
		return d.Close(closed)
	}

	return false
}

func (dd *Doors) PressButton(door uint8, duration time.Duration) (pressed bool, reason uint8) {
	if d, ok := dd.doors[door]; ok {
		return d.PressButton(duration)
	}

	return false, 0
}

func (dd *Doors) ControlState(door uint8) uint8 {
	if d, ok := dd.doors[door]; ok {
		return d.ControlState
	}

	return 0
}

func (dd *Doors) Delay(door uint8) Delay {
	if d, ok := dd.doors[door]; ok {
		return d.Delay
	}

	return 0
}

func (dd *Doors) EnableButton(door uint8, enabled bool) bool {
	if d, ok := dd.doors[door]; ok {
		return d.EnableButton(enabled)
	}

	return false
}

func (dd *Doors) IsOpen(door uint8) bool {
	if d, ok := dd.doors[door]; ok {
		return d.IsOpen()
	}

	return false
}

func (dd *Doors) IsButtonPressed(door uint8) bool {
	if d, ok := dd.doors[door]; ok {
		return d.IsButtonPressed()
	}

	return false
}

func (dd *Doors) IsUnlocked(door uint8) bool {
	if d, ok := dd.doors[door]; ok {
		return d.IsUnlocked()
	}

	return false
}

func (dd *Doors) IsProfileDisabled(door uint8) bool {
	if d, ok := dd.doors[door]; ok {
		return d.IsProfileDisabled()
	}

	return false
}

func (dd *Doors) IsNormallyClosed(door uint8) bool {
	if d, ok := dd.doors[door]; ok {
		return d.IsNormallyClosed()
	}

	return false
}

func (d Doors) MarshalJSON() ([]byte, error) {
	serializable := struct {
		Doors struct {
			Interlock uint8 `json:"interlock"`
			Door1     *Door `json:"1"`
			Door2     *Door `json:"2"`
			Door3     *Door `json:"3"`
			Door4     *Door `json:"4"`
		} `json:"doors"`
	}{
		Doors: struct {
			Interlock uint8 `json:"interlock"`
			Door1     *Door `json:"1"`
			Door2     *Door `json:"2"`
			Door3     *Door `json:"3"`
			Door4     *Door `json:"4"`
		}{
			Interlock: d.Interlock,
			Door1:     d.doors[1],
			Door2:     d.doors[2],
			Door3:     d.doors[3],
			Door4:     d.doors[4],
		},
	}

	return json.Marshal(serializable)
}

func (d *Doors) UnmarshalJSON(bytes []byte) error {
	serializable := struct {
		Doors struct {
			Interlock uint8 `json:"interlock"`
			Door1     *Door `json:"1"`
			Door2     *Door `json:"2"`
			Door3     *Door `json:"3"`
			Door4     *Door `json:"4"`
		} `json:"doors"`
	}{}

	if err := json.Unmarshal(bytes, &serializable); err != nil {
		return err
	}

	d.Interlock = serializable.Doors.Interlock
	d.doors[1] = serializable.Doors.Door1
	d.doors[2] = serializable.Doors.Door2
	d.doors[3] = serializable.Doors.Door3
	d.doors[4] = serializable.Doors.Door4

	for _, ix := range []uint8{1, 2, 3, 4} {
		if d.doors[ix].ControlState < 1 || d.doors[ix].ControlState > 3 {
			d.doors[ix].ControlState = 3
		}
	}

	return nil
}
