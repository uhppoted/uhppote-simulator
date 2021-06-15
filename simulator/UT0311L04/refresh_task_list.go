package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) refreshTaskList(addr *net.UDPAddr, request *messages.RefreshTaskListRequest) {
	if s.SerialNumber == request.SerialNumber {
		refreshed := false

		if request.MagicWord == 0x55aaaa55 {
			refreshed = s.TaskList.Refresh()
		}

		response := messages.RefreshTaskListResponse{
			SerialNumber: s.SerialNumber,
			Refreshed:    refreshed,
		}

		s.send(addr, &response)

		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
