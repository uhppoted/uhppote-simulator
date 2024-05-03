package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) setAddress(request *messages.SetAddressRequest) (any, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	if request.MagicWord != 0x55aaaa55 {
		fmt.Printf("ERROR: Invalid 'magic word' - expected: %08x, received:%08x", 0x55aaaa55, request.MagicWord)
		return nil, nil
	}

	s.IpAddress = request.Address
	s.SubnetMask = request.Mask
	s.Gateway = request.Gateway

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return nil, nil
}
