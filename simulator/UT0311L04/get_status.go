package UT0311L04

import (
	"net"
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) getStatus(addr *net.UDPAddr, request *messages.GetStatusRequest) {
	if s.SerialNumber == request.SerialNumber {
		utc := time.Now().UTC()
		datetime := utc.Add(time.Duration(s.TimeOffset))
		event := s.Events.Get(s.Events.Last)

		response := messages.GetStatusResponse{
			SerialNumber: s.SerialNumber,
			EventIndex:   s.Events.Last,
			SystemError:  s.SystemError,
			SystemDate:   types.SystemDate(datetime),
			SystemTime:   types.SystemTime(datetime),
			SequenceId:   s.SequenceId,
			SpecialInfo:  s.SpecialInfo,
			RelayState:   s.relays(),
			InputState:   s.InputState,

			Door1State: s.Doors[1].IsOpen(),
			Door2State: s.Doors[2].IsOpen(),
			Door3State: s.Doors[3].IsOpen(),
			Door4State: s.Doors[4].IsOpen(),

			Door1Button: s.Doors[1].IsButtonPressed(),
			Door2Button: s.Doors[2].IsButtonPressed(),
			Door3Button: s.Doors[3].IsButtonPressed(),
			Door4Button: s.Doors[4].IsButtonPressed(),
		}

		if event != nil {
			timestamp := event.Timestamp

			response.EventType = event.Type
			response.Granted = event.Granted
			response.Door = event.Door
			response.Direction = event.Direction
			response.CardNumber = event.Card
			response.Timestamp = &timestamp
			response.Reason = event.Reason
		}

		s.send(addr, &response)
	}
}
