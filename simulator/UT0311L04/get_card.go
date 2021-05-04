package UT0311L04

import (
	"net"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) getCardByID(addr *net.UDPAddr, request *messages.GetCardByIDRequest) {
	if request.SerialNumber == s.SerialNumber {
		response := messages.GetCardByIDResponse{
			SerialNumber: s.SerialNumber,
			CardNumber:   0,
		}

		for _, card := range s.Cards {
			if card != nil && request.CardNumber == card.CardNumber {
				response.CardNumber = card.CardNumber
				response.From = card.From
				response.To = card.To
				response.Door1 = card.Doors[1]
				response.Door2 = card.Doors[2]
				response.Door3 = card.Doors[3]
				response.Door4 = card.Doors[4]
				break
			}
		}

		s.send(addr, &response)
	}
}

func (s *UT0311L04) getCardByIndex(addr *net.UDPAddr, request *messages.GetCardByIndexRequest) {
	if request.SerialNumber == s.SerialNumber {
		response := messages.GetCardByIndexResponse{
			SerialNumber: s.SerialNumber,
			CardNumber:   0,
		}

		if request.Index > 0 && request.Index <= uint32(len(s.Cards)) {
			card := s.Cards[request.Index-1]

			if card != nil {
				response.CardNumber = card.CardNumber
				response.From = card.From
				response.To = card.To
				response.Door1 = card.Doors[1]
				response.Door2 = card.Doors[2]
				response.Door3 = card.Doors[3]
				response.Door4 = card.Doors[4]
			}

			s.send(addr, &response)
		}
	}
}
