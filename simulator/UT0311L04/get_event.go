package UT0311L04

import (
	"net"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) getEvent(addr *net.UDPAddr, request *messages.GetEventRequest) {
	if s.SerialNumber != request.SerialNumber {
		return
	}

	index := request.Index
	event := s.Events.Get(index)

	var timestamp types.DateTime

	if event.Timestamp != nil {
		timestamp = *event.Timestamp
	} else {
		timestamp = types.DateTime{}
	}

	response := messages.GetEventResponse{
		SerialNumber: s.SerialNumber,
		Index:        event.Index,
		Type:         event.Type,
		Granted:      event.Granted,
		Door:         event.Door,
		Direction:    event.Direction,
		CardNumber:   event.Card,
		Timestamp:    timestamp,
		Reason:       event.Reason,
	}

	s.send(addr, &response)
}
