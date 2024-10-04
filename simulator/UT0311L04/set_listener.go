package UT0311L04

import (
	"fmt"
	"time"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) setListener(request *messages.SetListenerRequest) (*messages.SetListenerResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	s.Listener = request.AddrPort
	s.AutoSend = request.Interval
	s.autosent = time.Now()

	response := messages.SetListenerResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    true,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
