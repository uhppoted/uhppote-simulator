package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func (s *UT0311L04) activateKeypads(addr *net.UDPAddr, request *messages.ActivateAccessKeypadsRequest) {
	set := func(keypad uint8, activated bool) {
		if activated {
			s.Keypads[keypad] = entities.KeypadBoth
		} else {
			s.Keypads[keypad] = entities.KeypadNone
		}
	}

	if request.SerialNumber == s.SerialNumber {
		set(1, request.Reader1)
		set(2, request.Reader2)
		set(3, request.Reader3)
		set(4, request.Reader4)

		response := messages.ActivateAccessKeypadsResponse{
			SerialNumber: s.SerialNumber,
			Succeeded:    true,
		}

		s.send(addr, &response)

		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
