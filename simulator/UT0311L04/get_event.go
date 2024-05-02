package UT0311L04

import (
	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) getEvent(request *messages.GetEventRequest) (*messages.GetEventResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	index := request.Index
	event := s.Events.Get(index)

	var timestamp types.DateTime

	if event.Timestamp.IsZero() {
		timestamp = types.DateTime{}
	} else {
		timestamp = event.Timestamp
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

	return &response, nil
}
