package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) setDoorPasscodes(addr *net.UDPAddr, request *messages.SetDoorPasscodesRequest) {
	if request.SerialNumber == s.SerialNumber {
		response := messages.SetDoorPasscodesResponse{
			SerialNumber: s.SerialNumber,
			Succeeded:    false,
		}

		door := request.Door

		if !(door < 1 || door > 4) {
			s.Doors.SetPasscodes(door, request.Passcode1, request.Passcode2, request.Passcode3, request.Passcode4)
			response.Succeeded = true
		}

		s.send(addr, &response)
		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
