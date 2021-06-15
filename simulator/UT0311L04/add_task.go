package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) addTask(addr *net.UDPAddr, request *messages.AddTaskRequest) {
	if s.SerialNumber == request.SerialNumber {
		task := types.Task{}

		added := s.TaskList.Add(task)

		response := messages.AddTaskResponse{
			SerialNumber: s.SerialNumber,
			Succeeded:    added,
		}

		s.send(addr, &response)

		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
