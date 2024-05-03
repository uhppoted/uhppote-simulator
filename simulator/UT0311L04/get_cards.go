package UT0311L04

import (
	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) getCards(request *messages.GetCardsRequest) (*messages.GetCardsResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	response := messages.GetCardsResponse{
		SerialNumber: s.SerialNumber,
		Records:      s.Cards.Size(),
	}

	return &response, nil
}
