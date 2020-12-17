package simulator

import (
	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
	"net"
)

type Simulator interface {
	DeviceID() uint32
	DeviceType() string
	FilePath() string
	SetTxQ(chan entities.Message)

	Handle(*net.UDPAddr, messages.Request)
	Save() error
	Delete() error

	Swipe(deviceID uint32, cardNumber uint32, door uint8) (bool, uint32)
	Open(deviceID uint32, door uint8) (uint32, error)
	Close(deviceID uint32, door uint8) (uint32, error)
}
