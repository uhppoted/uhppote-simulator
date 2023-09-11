package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) setSuperPasswords(addr *net.UDPAddr, request *messages.SetSuperPasswordsRequest) {
	if request.SerialNumber == s.SerialNumber {
		response := messages.SetSuperPasswordsResponse{
			SerialNumber: s.SerialNumber,
			Succeeded:    false,
		}

		door := request.Door

		if !(door < 1 || door > 4) {
			s.Doors.SetPasscodes(door, request.Password1, request.Password2, request.Password3, request.Password4)
			response.Succeeded = true
		}

		s.send(addr, &response)
		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
