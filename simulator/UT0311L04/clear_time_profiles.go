package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) clearTimeProfiles(addr *net.UDPAddr, request *messages.ClearTimeProfilesRequest) {
	if s.SerialNumber == request.SerialNumber {
		cleared := false

		if request.MagicWord == 0x55aaaa55 {
			cleared = s.TimeProfiles.Clear()
		}

		response := messages.ClearTimeProfilesResponse{
			SerialNumber: s.SerialNumber,
			Succeeded:    cleared,
		}

		s.send(addr, &response)

		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
