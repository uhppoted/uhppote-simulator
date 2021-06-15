package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) clearTaskList(addr *net.UDPAddr, request *messages.ClearTaskListRequest) {
	if s.SerialNumber == request.SerialNumber {
		cleared := false

		if request.MagicWord == 0x55aaaa55 {
			cleared = s.TaskList.Clear()
		}

		response := messages.ClearTaskListResponse{
			SerialNumber: s.SerialNumber,
			Succeeded:    cleared,
		}

		s.send(addr, &response)

		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
