package entities

import (
	"encoding/json"
	"slices"
	"sync"

	"github.com/uhppoted/uhppote-core/types"
)

var rules = map[types.AntiPassback]map[uint8]struct {
	deny  []uint8
	allow []uint8
}{
	types.Disabled: {
		1: {[]uint8{}, []uint8{}},
		2: {[]uint8{}, []uint8{}},
		3: {[]uint8{}, []uint8{}},
		4: {[]uint8{}, []uint8{}},
	},

	types.Readers12_34: {
		1: {[]uint8{1}, []uint8{2}},
		2: {[]uint8{2}, []uint8{1}},
		3: {[]uint8{3}, []uint8{4}},
		4: {[]uint8{4}, []uint8{3}},
	},

	types.Readers13_24: {
		1: {[]uint8{1}, []uint8{3}},
		2: {[]uint8{2}, []uint8{4}},
		3: {[]uint8{3}, []uint8{1}},
		4: {[]uint8{4}, []uint8{2}},
	},

	types.Readers1_23: {
		1: {[]uint8{1}, []uint8{2, 3}},
		2: {[]uint8{2, 3}, []uint8{1}},
		3: {[]uint8{2, 3}, []uint8{1}},
		4: {[]uint8{}, []uint8{}},
	},

	types.Readers1_234: {
		1: {[]uint8{1}, []uint8{2, 3, 4}},
		2: {[]uint8{2, 3, 4}, []uint8{1}},
		3: {[]uint8{2, 3, 4}, []uint8{1}},
		4: {[]uint8{2, 3, 4}, []uint8{1}},
	},
}

type pair struct {
	card uint32
	door uint8
}

type AntiPassback struct {
	antipassback types.AntiPassback
	deny         []pair
	guard        sync.RWMutex
}

func MakeAntiPassback(antipassback types.AntiPassback) *AntiPassback {
	return &AntiPassback{
		antipassback: antipassback,
	}
}

func (a *AntiPassback) Get() uint8 {
	a.guard.RLock()
	defer a.guard.RUnlock()

	u8 := uint8(a.antipassback)

	return u8
}

func (a *AntiPassback) Set(antipassback uint8) bool {
	a.guard.Lock()
	defer a.guard.Unlock()

	if antipassback <= 0x04 {
		a.antipassback = types.AntiPassback(antipassback)
		a.deny = []pair{}
		return true
	}

	return false
}

func (a *AntiPassback) Allow(card uint32, door uint8) bool {
	a.guard.RLock()
	defer a.guard.RUnlock()

	for _, v := range a.deny {
		if v.card == card && v.door == door {
			return false
		}
	}

	return true
}

func (a *AntiPassback) Allowed(card uint32, door uint8) {
	a.guard.Lock()
	defer a.guard.Unlock()

	a.append(card, rules[a.antipassback][door].deny...)
	a.delete(card, rules[a.antipassback][door].allow...)
}

func (a *AntiPassback) append(card uint32, doors ...uint8) {
	for _, door := range doors {
		f := func(v pair) bool {
			return v.card == card && v.door == door
		}

		if !slices.ContainsFunc(a.deny, f) {
			a.deny = append(a.deny, pair{
				card: card,
				door: door,
			})
		}
	}
}

func (a *AntiPassback) delete(card uint32, doors ...uint8) {
	f := func(v pair) bool {
		return v.card == card && slices.Contains(doors, v.door)
	}

	a.deny = slices.DeleteFunc(a.deny, f)
}

func (a *AntiPassback) MarshalJSON() ([]byte, error) {
	a.guard.RLock()
	defer a.guard.RUnlock()

	serializable := a.antipassback

	return json.Marshal(serializable)
}

func (a *AntiPassback) UnmarshalJSON(bytes []byte) error {
	a.guard.Lock()
	defer a.guard.Unlock()

	serializable := types.Disabled

	if err := json.Unmarshal(bytes, &serializable); err != nil {
		return err
	}

	a.antipassback = serializable

	return nil
}
