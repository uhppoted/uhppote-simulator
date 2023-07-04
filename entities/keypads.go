package entities

type Keypads map[uint8]bool

func MakeKeypads() Keypads {
	return Keypads{
		1: false,
		2: false,
		3: false,
		4: false,
	}
}
