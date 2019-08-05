package simulator

import (
	"net"
	"time"
	"uhppote/messages"
	"uhppote/types"
)

func (s *Simulator) getTime(addr *net.UDPAddr, request *messages.GetTimeRequest) {
	if s.SerialNumber == request.SerialNumber {

		utc := time.Now().UTC()
		datetime := utc.Add(time.Duration(s.TimeOffset))

		response := messages.GetTimeResponse{
			SerialNumber: s.SerialNumber,
			DateTime:     types.DateTime(datetime),
		}

		s.send(addr, &response)
	}
}
