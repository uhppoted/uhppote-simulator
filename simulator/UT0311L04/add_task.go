package UT0311L04

import (
	"fmt"
	"net"
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) addTask(addr *net.UDPAddr, request *messages.AddTaskRequest) {
	if s.SerialNumber == request.SerialNumber {
		task := types.Task{
			Task: types.TaskType(request.Task),
			Door: request.Door,
			From: request.From,
			To:   request.To,
			Weekdays: types.Weekdays{
				time.Monday:    request.Monday,
				time.Tuesday:   request.Tuesday,
				time.Wednesday: request.Wednesday,
				time.Thursday:  request.Thursday,
				time.Friday:    request.Friday,
				time.Saturday:  request.Saturday,
				time.Sunday:    request.Sunday,
			},
			Start: request.Start,
			Cards: request.MoreCards,
		}

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
