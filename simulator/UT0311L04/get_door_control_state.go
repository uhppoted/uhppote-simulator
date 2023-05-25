package UT0311L04

import (
	"github.com/uhppoted/uhppote-core/messages"
	"net"
)

func (s *UT0311L04) getDoorControlState(addr *net.UDPAddr, request *messages.GetDoorControlStateRequest) {
	if request.SerialNumber == s.SerialNumber {

		if !(request.Door < 1 || request.Door > 4) {
			response := messages.GetDoorControlStateResponse{
				SerialNumber: s.SerialNumber,
				Door:         request.Door,
				ControlState: s.Doors.ControlState(request.Door),
				Delay:        s.Doors.Delay(request.Door).Seconds(),
			}

			s.send(addr, &response)
		}
	}
}
