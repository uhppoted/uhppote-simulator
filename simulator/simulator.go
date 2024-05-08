package simulator

import (
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
)

type Simulator interface {
	DeviceID() uint32
	DeviceType() string
	FilePath() string

	Handle(messages.Request) (any, error)
	RunTasks()
	Save() error
	Delete() error

	Swipe(cardNumber uint32, door uint8, direction entities.Direction, PIN uint32) (bool, error)
	Passcode(door uint8, passcode uint32) (bool, error)
	Open(door uint8) (bool, error)
	Close(door uint8) (bool, error)
	ButtonPressed(door uint8, duration time.Duration) (bool, error)
}
