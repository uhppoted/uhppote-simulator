package UT0311L04

import (
	"net"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) getListener(request *messages.GetListenerRequest) (*messages.GetListenerResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	address := net.IPv4(0, 0, 0, 0)
	port := uint16(0)

	if s.Listener.IsValid() {
		addr := s.Listener.Addr().As4()
		address = net.IPv4(addr[0], addr[1], addr[2], addr[3])
		port = s.Listener.Port()
	}

	response := messages.GetListenerResponse{
		SerialNumber: s.SerialNumber,
		Address:      address,
		Port:         port,
	}

	return &response, nil
}
