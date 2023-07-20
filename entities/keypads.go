package entities

import (
	"encoding/json"

	"github.com/uhppoted/uhppote-simulator/log"
)

type Keypad uint8
type Keypads map[uint8]Keypad

const (
	KeypadNone Keypad = 0
	KeypadIn   Keypad = 1
	KeypadOut  Keypad = 2
	KeypadBoth Keypad = 3
)

func MakeKeypads() Keypads {
	return Keypads{
		1: KeypadNone,
		2: KeypadNone,
		3: KeypadNone,
		4: KeypadNone,
	}
}

func (k *Keypads) UnmarshalJSON(bytes []byte) error {
	keypads := map[uint8]Keypad{
		1: KeypadNone,
		2: KeypadNone,
		3: KeypadNone,
		4: KeypadNone,
	}

	if k != nil {
		keypads[1] = (*k)[1]
		keypads[2] = (*k)[2]
		keypads[3] = (*k)[3]
		keypads[4] = (*k)[4]
	}

	if err := json.Unmarshal(bytes, &keypads); err != nil {
		log.Warnf("%v", err)
	} else {
		*k = Keypads{
			1: keypads[1],
			2: keypads[2],
			3: keypads[3],
			4: keypads[4],
		}
	}

	return nil
}
