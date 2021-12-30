package UT0311L04

import (
	"github.com/uhppoted/uhppote-core/messages"
	"net"
)

func (s *UT0311L04) getEvent(addr *net.UDPAddr, request *messages.GetEventRequest) {
	if s.SerialNumber != request.SerialNumber {
		return
	}

	index := request.Index

	event := s.Events.Get(index)
	response := messages.GetEventResponse{
		SerialNumber: s.SerialNumber,
		Index:        event.Index,
		Type:         event.Type,
		Granted:      event.Granted,
		Door:         event.Door,
		Direction:    event.Direction,
		CardNumber:   event.Card,
		Timestamp:    event.Timestamp,
		Reason:       event.Reason,
	}

	s.send(addr, &response)
}
