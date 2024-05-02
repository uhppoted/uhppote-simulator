package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func (s *UT0311L04) activateKeypads(addr *net.UDPAddr, request *messages.ActivateAccessKeypadsRequest) (*messages.ActivateAccessKeypadsResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	set := func(keypad uint8, activated bool) {
		if activated {
			s.Keypads[keypad] = entities.KeypadBoth
		} else {
			s.Keypads[keypad] = entities.KeypadNone
		}
	}

	set(1, request.Reader1)
	set(2, request.Reader2)
	set(3, request.Reader3)
	set(4, request.Reader4)

	response := messages.ActivateAccessKeypadsResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    true,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
