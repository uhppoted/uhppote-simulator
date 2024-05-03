package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) refreshTaskList(request *messages.RefreshTaskListRequest) (*messages.RefreshTaskListResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	refreshed := false

	if request.MagicWord == 0x55aaaa55 {
		refreshed = s.TaskList.Refresh()
	}

	response := messages.RefreshTaskListResponse{
		SerialNumber: s.SerialNumber,
		Refreshed:    refreshed,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
