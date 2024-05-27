package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) setListener(request *messages.SetListenerRequest) (*messages.SetListenerResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	listener := request.AddrPort
	s.Listener = listener

	response := messages.SetListenerResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    true,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
