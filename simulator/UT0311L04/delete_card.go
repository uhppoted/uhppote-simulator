package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) deleteCard(request *messages.DeleteCardRequest) (*messages.DeleteCardResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	deleted := s.Cards.Delete(request.CardNumber)

	response := messages.DeleteCardResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    deleted,
	}

	if deleted {
		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}

	return &response, nil
}
