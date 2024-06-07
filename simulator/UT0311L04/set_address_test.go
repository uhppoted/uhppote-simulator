package UT0311L04

import (
	"net"
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestHandleSetAddress(t *testing.T) {
	request := messages.SetAddressRequest{
		SerialNumber: 12345,
		Address:      net.IPv4(10, 0, 0, 100),
		Mask:         net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MagicWord:    0x55aaaa55,
	}

	testHandle(&request, nil, t)
}
