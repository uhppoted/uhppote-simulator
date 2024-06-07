package UT0311L04

import (
	"net/netip"
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
)

func TestHandleGetListener(t *testing.T) {
	request := messages.GetListenerRequest{
		SerialNumber: 12345,
	}

	response := messages.GetListenerResponse{
		SerialNumber: 12345,
		AddrPort:     netip.MustParseAddrPort("10.0.0.10:43210"),
	}

	testHandle(&request, &response, t)
}
