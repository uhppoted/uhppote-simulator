package UT0311L04

import (
	"fmt"
	"net"
	"net/netip"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func (s *UT0311L04) restoreDefaultParameters(request *messages.RestoreDefaultParametersRequest) (*messages.RestoreDefaultParametersResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	reset := false

	if request.MagicWord == 0x55aaaa55 {
		reset = true
		s.IpAddress = net.IPv4(0, 0, 0, 0)
		s.SubnetMask = net.IPv4(255, 255, 255, 0)
		s.Gateway = net.IPv4(0, 0, 0, 0)
		s.Listener = netip.AddrPort{}
		s.AntiPassback = entities.AntiPassback{}

		for _, door := range []uint8{1, 2, 3, 4} {
			s.Doors.SetControlState(door, entities.Controlled)
			s.Doors.SetDelay(door, entities.DelayFromSeconds(5))
			s.Doors.SetPasscodes(door)
		}

		if !s.Events.Clear() {
			reset = false
		}

		if !s.Cards.DeleteAll() {
			reset = false
		}
	}

	response := messages.RestoreDefaultParametersResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    reset,
	}

	if reset {
		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}

	return &response, nil
}
