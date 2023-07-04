package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) activateKeypads(addr *net.UDPAddr, request *messages.ActivateAccessKeypadsRequest) {
	if request.SerialNumber == s.SerialNumber {
		s.Keypads[1] = request.Reader1
		s.Keypads[2] = request.Reader2
		s.Keypads[3] = request.Reader3
		s.Keypads[4] = request.Reader4

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
