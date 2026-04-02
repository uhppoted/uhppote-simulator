package entities

import (
	"encoding/json"
	"fmt"
	"iter"
	"strings"

	"github.com/uhppoted/uhppote-core/types"
)

type Card struct {
	CardNumber uint32              `json:"card"`
	From       types.Date          `json:"from"`
	To         types.Date          `json:"to"`
	Doors      map[uint8]uint8     `json:"doors,omitempty"`
	PIN        uint32              `json:"PIN,omitempty"`
	FirstCard  FirstCardPrivileges `json:"firstcard"`
}

type FirstCardPrivileges struct {
	Door1 bool
	Door2 bool
	Door3 bool
	Door4 bool
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

	return fmt.Errorf("insufficient space in card list")
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
	for ix := range *l {
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

func (f FirstCardPrivileges) ToUint8() uint8 {
	v := uint8(0)

	masks := map[uint8]uint8{
		1: 0x01,
		2: 0x02,
		3: 0x04,
		4: 0x08,
	}

	for d, enabled := range f.ForEach() {
		if enabled {
			v = v | masks[d]
		}
	}

	return v
}

func (f FirstCardPrivileges) String() string {
	var v []string

	for d, enabled := range f.ForEach() {
		if enabled {
			v = append(v, fmt.Sprintf("%d", d))
		}
	}

	if len(v) == 0 {
		return "-"
	}

	return strings.Join(v, ",")
}

func (f FirstCardPrivileges) Has(door uint8) bool {
	for d, enabled := range f.ForEach() {
		if d == door {
			return enabled
		}
	}

	return false
}

func (f FirstCardPrivileges) ForEach() iter.Seq2[uint8, bool] {
	return func(yield func(uint8, bool) bool) {
		if !yield(1, f.Door1) {
			return
		}

		if !yield(2, f.Door2) {
			return
		}

		if !yield(3, f.Door3) {
			return
		}

		if !yield(4, f.Door4) {
			return
		}
	}
}

func (f FirstCardPrivileges) MarshalJSON() ([]byte, error) {
	m := map[uint8]bool{
		1: f.Door1,
		2: f.Door2,
		3: f.Door3,
		4: f.Door4,
	}

	return json.Marshal(m)
}

func (f *FirstCardPrivileges) UnmarshalJSON(bytes []byte) error {
	m := map[uint8]bool{
		1: false,
		2: false,
		3: false,
		4: false,
	}

	if err := json.Unmarshal(bytes, &m); err != nil {
		return err
	}

	f.Door1 = m[1]
	f.Door2 = m[2]
	f.Door3 = m[3]
	f.Door4 = m[4]

	return nil
}
