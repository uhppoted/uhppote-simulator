package UT0311L04

import (
	"codeberg.org/uhppoted/uhppoted-core/messages"
	"codeberg.org/uhppoted/uhppoted-core/types"

	"codeberg.org/uhppoted/uhppoted-simulator/entities"
)

func (s *UT0311L04) getCardByID(request *messages.GetCardByIDRequest) (*messages.GetCardByIDResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	f := func(card *entities.Card) bool {
		return card != nil && request.CardNumber == card.CardNumber
	}

	if _, card := s.Cards.Find(f); card != nil {
		return &messages.GetCardByIDResponse{
			SerialNumber: s.SerialNumber,
			CardNumber:   card.CardNumber,
			From:         card.From,
			To:           card.To,
			Door1:        card.Doors[1],
			Door2:        card.Doors[2],
			Door3:        card.Doors[3],
			Door4:        card.Doors[4],
			PIN:          pin(*card),
			FirstCard:    card.FirstCard.ToUint8(),
		}, nil
	}

	return &messages.GetCardByIDResponse{
		SerialNumber: s.SerialNumber,
		CardNumber:   0,
	}, nil
}

func (s *UT0311L04) getCardByIndex(request *messages.GetCardByIndexRequest) (*messages.GetCardByIndexResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	ix, card := s.Cards.Get(request.Index)
	if ix != -1 && card != nil {
		return &messages.GetCardByIndexResponse{
			SerialNumber: s.SerialNumber,
			CardNumber:   card.CardNumber,
			From:         card.From,
			To:           card.To,
			Door1:        card.Doors[1],
			Door2:        card.Doors[2],
			Door3:        card.Doors[3],
			Door4:        card.Doors[4],
			PIN:          pin(*card),
			FirstCard:    card.FirstCard.ToUint8(),
		}, nil
	}

	if ix != -1 && card == nil {
		return &messages.GetCardByIndexResponse{
			SerialNumber: s.SerialNumber,
			CardNumber:   0,
		}, nil
	}

	return nil, nil
}

func pin(card entities.Card) types.PIN {
	if card.PIN < 1000000 {
		return types.PIN(card.PIN)
	}

	return types.PIN(0)
}
