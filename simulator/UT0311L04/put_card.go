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

	response := messages.PutCardResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    false,
	}

	if !request.From.IsZero() && !request.To.IsZero() {
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

		if err := s.Cards.Put(&card); err != nil {
			fmt.Printf("WARN:  %v\n", err)
			response.Succeeded = false
		} else {
			response.Succeeded = true
		}

		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}

	return &response, nil
}
