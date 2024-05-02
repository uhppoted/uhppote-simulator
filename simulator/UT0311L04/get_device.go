package UT0311L04

import (
	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) getDevice(request *messages.GetDeviceRequest) (*messages.GetDeviceResponse, error) {
	if request.SerialNumber != 0 && request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	response := messages.GetDeviceResponse{
		SerialNumber: s.SerialNumber,
		IpAddress:    s.IpAddress,
		SubnetMask:   s.SubnetMask,
		Gateway:      s.Gateway,
		MacAddress:   s.MacAddress,
		Version:      types.Version(s.Version),
		Date:         types.Date(*s.Released),
	}

	return &response, nil
}
