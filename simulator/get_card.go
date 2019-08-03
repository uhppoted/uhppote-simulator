package simulator

import (
	"uhppote/messages"
)

func (s *Simulator) GetCardById(request *messages.GetCardByIdRequest) *messages.GetCardByIdResponse {
	if request.SerialNumber != s.SerialNumber {
		return nil
	}

	response := messages.GetCardByIdResponse{
		SerialNumber: s.SerialNumber,
	}

	for _, card := range s.Cards {
		if request.CardNumber == card.CardNumber {
			response.CardNumber = card.CardNumber
			response.From = &card.From
			response.To = &card.To
			response.Door1 = card.Doors[1]
			response.Door2 = card.Doors[2]
			response.Door3 = card.Doors[3]
			response.Door4 = card.Doors[4]
			break
		}
	}

	return &response
}

func (s *Simulator) getCardByIndex(request *messages.GetCardByIndexRequest) *messages.GetCardByIndexResponse {
	if request.SerialNumber != s.SerialNumber {
		return nil
	}

	response := messages.GetCardByIndexResponse{
		SerialNumber: s.SerialNumber,
	}

	if request.Index > 0 && request.Index <= uint32(len(s.Cards)) {
		card := s.Cards[request.Index-1]
		response.CardNumber = card.CardNumber
		response.From = &card.From
		response.To = &card.To
		response.Door1 = card.Doors[1]
		response.Door2 = card.Doors[2]
		response.Door3 = card.Doors[3]
		response.Door4 = card.Doors[4]
	}

	return &response
}
