package entities

import (
	"encoding/json"
	"fmt"

	"github.com/uhppoted/uhppote-core/types"
)

type Card struct {
	CardNumber uint32          `json:"card"`
	From       *types.Date     `json:"from,omitempty"`
	To         *types.Date     `json:"to,omitempty"`
	Doors      map[uint8]uint8 `json:"doors,omitempty"`
}

type CardList [64]*Card

func (l *CardList) Size() uint32 {
	var count uint32 = 0

	for _, card := range *l {
		if card != nil && card.CardNumber != 0 && card.CardNumber != 0xffffffff {
			count++
		}
	}

	return count
}

func (l *CardList) Put(card *Card) error {
	if card == nil {
		return nil
	}

	for ix, c := range *l {
		if c != nil && c.CardNumber == card.CardNumber {
			l[ix] = card
			return nil
		}
	}

	for ix, c := range *l {
		if c == nil || c.CardNumber == 0 || c.CardNumber == 0xffffffff {
			l[ix] = card
			return nil
		}
	}

	return fmt.Errorf("Insufficient space in card list")
}

func (l *CardList) Delete(cardNumber uint32) bool {
	if cardNumber != 0 && cardNumber != 0xffffffff {
		for ix, c := range *l {
			if c != nil && c.CardNumber == cardNumber {
				(*l)[ix] = &Card{
					CardNumber: 0xffffffff,
				}

				return true
			}
		}
	}

	return false
}

func (l *CardList) DeleteAll() bool {
	for ix, _ := range *l {
		(*l)[ix] = nil
	}

	return true
}

func (l CardList) MarshalJSON() ([]byte, error) {
	N := 0
	for i, c := range l {
		if c != nil {
			N = i + 1
		}
	}

	list := make([]*Card, N)

	copy(list, l[0:N])

	return json.Marshal(list)
}
