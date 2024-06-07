package UT0311L04

import (
	"net/netip"
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestHandleSetListener(t *testing.T) {
	request := messages.SetListenerRequest{
		SerialNumber: 12345,
		AddrPort:     netip.MustParseAddrPort("10.0.0.1:43210"),
	}

	response := messages.SetListenerResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}
