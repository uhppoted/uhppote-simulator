package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) setPCControl(addr *net.UDPAddr, request *messages.SetPCControlRequest) {
	if request.SerialNumber == s.SerialNumber {
		ok := false

		if request.MagicWord == 0x55aaaa55 {
			s.PCControl = request.Enable
			ok = true
		}

		response := messages.SetPCControlResponse{
			SerialNumber: s.SerialNumber,
			Succeeded:    ok,
		}

		s.send(addr, &response)

		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
