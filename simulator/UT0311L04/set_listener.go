package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) setListener(request *messages.SetListenerRequest) (*messages.SetListenerResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	listener := net.UDPAddr{IP: request.Address, Port: int(request.Port)}
	s.Listener = &listener

	response := messages.SetListenerResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    true,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
