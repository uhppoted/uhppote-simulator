package UT0311L04

import (
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) getStatus(request *messages.GetStatusRequest) (*messages.GetStatusResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	utc := time.Now().UTC()
	datetime := utc.Add(time.Duration(s.TimeOffset))
	event := s.Events.Get(0xffffffff)

	response := messages.GetStatusResponse{
		SerialNumber: s.SerialNumber,
		EventIndex:   0,
		SystemError:  s.SystemError,
		SystemDate:   types.SystemDate(datetime),
		SystemTime:   types.SystemTime(datetime),
		SequenceId:   s.SequenceId,
		SpecialInfo:  s.SpecialInfo,
		RelayState:   s.relays(),
		InputState:   s.InputState,

		Door1State: s.Doors.IsOpen(1),
		Door2State: s.Doors.IsOpen(2),
		Door3State: s.Doors.IsOpen(3),
		Door4State: s.Doors.IsOpen(4),

		Door1Button: s.Doors.IsButtonPressed(1),
		Door2Button: s.Doors.IsButtonPressed(2),
		Door3Button: s.Doors.IsButtonPressed(3),
		Door4Button: s.Doors.IsButtonPressed(4),
	}

	response.EventIndex = event.Index
	response.EventType = event.Type
	response.Granted = event.Granted
	response.Door = event.Door
	response.Direction = event.Direction
	response.CardNumber = event.Card
	response.Timestamp = event.Timestamp
	response.Reason = event.Reason

	return &response, nil
}
