package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) deleteCards(request *messages.DeleteCardsRequest) (*messages.DeleteCardsResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	deleted := false

	if request.MagicWord == 0x55aaaa55 {
		deleted = s.Cards.DeleteAll()
	}

	response := messages.DeleteCardsResponse{
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
