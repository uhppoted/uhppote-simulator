package entities

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
