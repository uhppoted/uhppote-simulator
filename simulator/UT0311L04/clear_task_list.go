package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) clearTaskList(request *messages.ClearTaskListRequest) (*messages.ClearTaskListResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	cleared := false

	if request.MagicWord == 0x55aaaa55 {
		cleared = s.TaskList.Clear()
	}

	response := messages.ClearTaskListResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    cleared,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
