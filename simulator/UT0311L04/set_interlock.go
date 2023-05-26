package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) setInterlock(addr *net.UDPAddr, request *messages.SetInterlockRequest) {
	if request.SerialNumber == s.SerialNumber {
		ok := false

		if request.Interlock <= 4 {
			s.Doors.Interlock = request.Interlock
			ok = true
		}

		response := messages.SetInterlockResponse{
			SerialNumber: s.SerialNumber,
			Succeeded:    ok,
		}

		s.send(addr, &response)

		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
