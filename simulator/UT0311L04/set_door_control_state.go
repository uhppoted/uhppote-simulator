package UT0311L04

import (
	"fmt"
	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
	"net"
)

func (s *UT0311L04) setDoorControlState(addr *net.UDPAddr, request *messages.SetDoorControlStateRequest) {
	if request.SerialNumber == s.SerialNumber {
		door := request.Door
		if door < 1 || door > 4 {
			fmt.Printf("ERROR: Invalid door' - expected: [1..4], received:%d", request.Door)
			return
		}

		s.Doors.SetControlState(door, request.ControlState)
		s.Doors.SetDelay(door, entities.Delay(uint64(request.Delay)*1000000000))

		response := messages.SetDoorControlStateResponse{
			SerialNumber: s.SerialNumber,
			Door:         door,
			ControlState: s.Doors.ControlState(door),
			Delay:        s.Doors.Delay(door).Seconds(),
		}

		s.send(addr, &response)

		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
