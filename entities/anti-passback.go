package entities

import (
	"encoding/json"

	"github.com/uhppoted/uhppote-core/types"
)

type AntiPassback struct {
	antipassback types.AntiPassback
}

func MakeAntiPassback(antipassback types.AntiPassback) AntiPassback {
	return AntiPassback{
		antipassback: antipassback,
	}
}

func (a AntiPassback) Get() uint8 {
	return uint8(a.antipassback)
}

func (a *AntiPassback) Set(antipassback uint8) bool {
	if antipassback <= 0x04 {
		a.antipassback = types.AntiPassback(antipassback)
		return true
	}

	return false
}

func (a AntiPassback) Deny(card uint32, door uint8) bool {
	switch a.antipassback {
	case types.Readers12_34:
		// TODO
		return true

	case types.Readers13_24:
		// TODO
		return true

	case types.Readers1_23:
		//TODO
		return true

	case types.Readers1_234:
		//TODO
		return true
	}

	return false
}

func (a AntiPassback) MarshalJSON() ([]byte, error) {
	serializable := a.antipassback

	return json.Marshal(serializable)
}

func (a *AntiPassback) UnmarshalJSON(bytes []byte) error {
	serializable := types.Disabled

	if err := json.Unmarshal(bytes, &serializable); err != nil {
		return err
	}

	a.antipassback = serializable

	return nil
}
