package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) recordSpecialEvents(addr *net.UDPAddr, request *messages.RecordSpecialEventsRequest) {
	if request.SerialNumber == s.SerialNumber {
		s.RecordSpecialEvents = request.Enable

		response := messages.RecordSpecialEventsResponse{
			SerialNumber: s.SerialNumber,
			Succeeded:    true,
		}

		s.send(addr, &response)

		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
