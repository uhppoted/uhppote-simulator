package UT0311L04

import (
	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) getCardByID(request *messages.GetCardByIDRequest) (*messages.GetCardByIDResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

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

			if card.PIN < 1000000 {
				response.PIN = types.PIN(card.PIN)
			}

			break
		}
	}

	return &response, nil
}

func (s *UT0311L04) getCardByIndex(request *messages.GetCardByIndexRequest) (*messages.GetCardByIndexResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

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

			if card.PIN < 1000000 {
				response.PIN = types.PIN(card.PIN)
			}
		}

		return &response, nil
	}

	return nil, nil
}
