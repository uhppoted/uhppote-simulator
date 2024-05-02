package entities

import (
	"net"
)

type Message struct {
	Destination *net.UDPAddr
	Message     any
	Event       bool
}
