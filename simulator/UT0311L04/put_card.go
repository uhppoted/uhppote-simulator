package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func (s *UT0311L04) putCard(request *messages.PutCardRequest) (*messages.PutCardResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	card := entities.Card{
		CardNumber: request.CardNumber,
		From:       request.From,
		To:         request.To,
		Doors: map[uint8]uint8{1: request.Door1,
			2: request.Door2,
			3: request.Door3,
			4: request.Door4,
		},
	}

	if request.PIN < 1000000 {
		card.PIN = uint32(request.PIN)
	}

	var succeeded bool
	if err := s.Cards.Put(&card); err != nil {
		fmt.Printf("WARN:  %v\n", err)
		succeeded = false
	} else {
		succeeded = true
	}

	response := messages.PutCardResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    succeeded,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
