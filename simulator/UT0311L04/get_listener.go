package UT0311L04

import (
	"net/netip"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) getListener(request *messages.GetListenerRequest) (*messages.GetListenerResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	addr := netip.MustParseAddrPort("0.0.0.0:0")
	if s.Listener.IsValid() {
		addr = s.Listener
	}

	response := messages.GetListenerResponse{
		SerialNumber: s.SerialNumber,
		AddrPort:     addr,
		Interval:     s.AutoSend,
	}

	return &response, nil
}
