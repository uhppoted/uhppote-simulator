package UT0311L04

import (
	"net"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) getDevice(addr *net.UDPAddr, request *messages.GetDeviceRequest) {
	if request.SerialNumber == 0 || request.SerialNumber == s.SerialNumber {

		response := messages.GetDeviceResponse{
			SerialNumber: s.SerialNumber,
			IpAddress:    s.IpAddress,
			SubnetMask:   s.SubnetMask,
			Gateway:      s.Gateway,
			MacAddress:   s.MacAddress,
			Version:      types.Version(s.Version),
			Date:         types.Date(*s.Released),
		}

		s.send(addr, &response)
	}
}
