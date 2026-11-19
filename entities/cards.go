package entities

import (
	"encoding/json"
	"fmt"
	"iter"
	"slices"
	"strings"

	"codeberg.org/uhppoted/uhppoted-core/types"
)

type CardList interface {
	Size() uint32
	Find(f func(*Card) bool) (int, *Card)
	Get(index uint32) (int, *Card)
	Put(card *Card) error
	Delete(cardNumber uint32) bool
	DeleteAll() bool
}

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

type cardList struct {
	cards []*Card
}

var cardsListSize uint = 64

func SetCardsListSize(N uint) {
	if N <= 20000 {
		cardsListSize = N
	}
}

func NewCardList(cards ...*Card) CardList {
	list := cardList{
		cards: make([]*Card, cardsListSize),
	}

	for ix, card := range cards {
		if ix < len(list.cards) {
			list.cards[ix] = card
		}
	}

	return &list
}

func (l *cardList) All() iter.Seq2[int, *Card] {
	return func(yield func(int, *Card) bool) {
		for i, card := range l.cards {
			if !yield(i, card) {
				return
			}
		}
	}
}

func (l *cardList) Size() uint32 {
	var count uint32 = 0

	for _, card := range l.All() {
		if card != nil && card.CardNumber != 0 && card.CardNumber != 0xffffffff {
			count++
		}
	}

	return count
}

func (l *cardList) Put(card *Card) error {
	if card == nil {
		return nil
	}

	for ix, c := range l.All() {
		if c != nil && c.CardNumber == card.CardNumber {
			l.cards[ix] = card
			return nil
		}
	}

	for ix, c := range l.All() {
		if c == nil || c.CardNumber == 0 || c.CardNumber == 0xffffffff {
			l.cards[ix] = card
			return nil
		}
	}

	return fmt.Errorf("insufficient space in card list")
}

func (l *cardList) Delete(cardNumber uint32) bool {
	if cardNumber != 0 && cardNumber != 0xffffffff {
		for ix, c := range l.All() {
			if c != nil && c.CardNumber == cardNumber {
				l.cards[ix] = &Card{
					CardNumber: 0xffffffff,
				}

				return true
			}
		}
	}

	return false
}

func (l *cardList) DeleteAll() bool {
	for ix := range l.All() {
		l.cards[ix] = nil
	}

	return true
}

func (l *cardList) Find(f func(*Card) bool) (int, *Card) {
	ix := slices.IndexFunc(l.cards[:], f)
	if ix == -1 {
		return -1, nil
	}
	return ix, l.cards[ix]
}

func (l cardList) Get(index uint32) (int, *Card) {
	if index > 0 && index <= uint32(len(l.cards)) {
		return int(index), l.cards[index-1]
	}

	return -1, nil
}

func (l *cardList) MarshalJSON() ([]byte, error) {
	N := 0
	for i, c := range l.All() {
		if c != nil {
			N = i + 1
		}
	}

	list := make([]*Card, N)

	copy(list, l.cards[0:N])

	return json.Marshal(list)
}

func (l *cardList) UnmarshalJSON(bytes []byte) error {
	var list []*Card
	if err := json.Unmarshal(bytes, &list); err != nil {
		return err
	}

	cards := make([]*Card, cardsListSize)
	copy(cards, list)
	l.cards = cards

	return nil
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
