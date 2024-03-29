package entities

import (
	"encoding/json"
	"fmt"
	"strings"
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

func (dd *Doors) SetPasscodes(door uint8, passcodes ...uint32) {
	if d, ok := dd.doors[door]; ok {
		p := []uint32{0, 0, 0, 0}

		for i := 0; i < 4; i++ {
			if i < len(passcodes) && passcodes[i] < 1000000 {
				p[i] = passcodes[i]
			}
		}

		d.Passcodes = p
	}
}

func (dd *Doors) PressButton(door uint8, duration time.Duration) (pressed bool, reason uint8) {
	if d, ok := dd.doors[door]; ok {
		return d.PressButton(duration)
	}

	return false, 0
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

func (dd *Doors) UnlockWithPasscode(door uint8, passcode uint32, duration time.Duration) bool {
	if d, ok := dd.doors[door]; ok {
		return d.UnlockWithPasscode(passcode, duration)
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

func (dd *Doors) Passcodes(door uint8) []uint32 {
	if d, ok := dd.doors[door]; ok {
		return d.Passcodes
	}

	return []uint32{}
}

func (dd *Doors) EnableButton(door uint8, enabled bool) bool {
	if d, ok := dd.doors[door]; ok {
		return d.EnableButton(enabled)
	}

	return false
}

func (dd *Doors) IsOpen(doors ...uint8) bool {
	for _, door := range doors {
		if d, ok := dd.doors[door]; ok && d.IsOpen() {
			return true
		}
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

func (dd *Doors) IsInterlocked(door uint8) bool {
	switch dd.Interlock {
	case 0:
		return false

	case 1:
		if door == 1 && dd.IsOpen(2) || door == 2 && dd.IsOpen(1) {
			return true
			// } else if door == 3 && dd.IsOpen(4) || door == 4 && dd.IsOpen(3) {
			// 	return true
		}

	case 2:
		if door == 3 && dd.IsOpen(4) || door == 4 && dd.IsOpen(3) {
			return true
		}

	case 3:
		if door == 1 && dd.IsOpen(2) || door == 2 && dd.IsOpen(1) {
			return true
		} else if door == 3 && dd.IsOpen(4) || door == 4 && dd.IsOpen(3) {
			return true
		}

	case 4:
		if door == 1 && (dd.IsOpen(2, 3)) {
			return true
		} else if door == 2 && (dd.IsOpen(1, 3)) {
			return true
		} else if door == 3 && (dd.IsOpen(1, 2)) {
			return true
		}

	case 8:
		if door == 1 && dd.IsOpen(2, 3, 4) {
			return true
		} else if door == 2 && dd.IsOpen(1, 3, 4) {
			return true
		} else if door == 3 && dd.IsOpen(1, 2, 4) {
			return true
		} else if door == 4 && dd.IsOpen(1, 2, 3) {
			return true
		}
	}

	return false
}

func (d Doors) String() string {
	var b strings.Builder

	for _, door := range []uint8{1, 2, 3, 4} {
		fmt.Fprintf(&b, "door:%v mode:%v, delay:%vs\n", door, d.doors[door].ControlState, d.doors[door].Delay)
	}

	return b.String()
}

func (d Doors) MarshalJSON() ([]byte, error) {
	serializable := struct {
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
	}

	return json.Marshal(serializable)
}

func (d *Doors) UnmarshalJSON(bytes []byte) error {
	serializable := struct {
		Interlock uint8 `json:"interlock"`
		Door1     *Door `json:"1"`
		Door2     *Door `json:"2"`
		Door3     *Door `json:"3"`
		Door4     *Door `json:"4"`
	}{}

	if err := json.Unmarshal(bytes, &serializable); err != nil {
		return err
	}

	d.Interlock = serializable.Interlock
	d.doors = map[uint8]*Door{
		1: NewDoor(1),
		2: NewDoor(2),
		3: NewDoor(3),
		4: NewDoor(4),
	}

	if serializable.Door1 != nil {
		d.doors[1] = serializable.Door1
	}

	if serializable.Door2 != nil {
		d.doors[2] = serializable.Door2
	}

	if serializable.Door3 != nil {
		d.doors[3] = serializable.Door3
	}

	if serializable.Door4 != nil {
		d.doors[4] = serializable.Door4
	}

	for _, ix := range []uint8{1, 2, 3, 4} {
		if d.doors[ix].ControlState < 1 || d.doors[ix].ControlState > 3 {
			d.doors[ix].ControlState = 3
		}
	}

	return nil
}
