package entities

import (
	"reflect"
	"testing"
	"time"

	"github.com/uhppoted/uhppote-core/types"
)

var date = func(s string) *types.Date {
	d, _ := time.ParseInLocation("2006-01-02", s, time.Local)
	p := types.Date(d)
	return &p
}

func TestCardListSize(t *testing.T) {
	cards := CardList{}
	expected := fill(&cards)

	if N := cards.Size(); N != expected {
		t.Errorf("CardList.Size returned %v, expected: %v", N, expected)
	}
}

func TestCardListPutWithNewCard(t *testing.T) {
	cards := CardList{}
	expected := CardList{}

	fill(&cards)
	fill(&expected)

	card := Card{
		CardNumber: uint32(5000019),
		From:       date("2020-02-03"),
		To:         date("2020-11-30"),
		Doors: map[uint8]uint8{
			1: 0,
			2: 1,
			3: 1,
			4: 0,
		},
	}

	expected[3] = &card

	err := cards.Put(&card)
	if err != nil {
		t.Fatalf("Unexpected error adding card to list: %v", err)
	}

	if !reflect.DeepEqual(cards, expected) {
		for i, c := range cards {
			if !reflect.DeepEqual(c, expected[i]) {
				t.Errorf("Invalid CardList entry %v\n   expected: %v\n   got:      %v", i, *expected[i], *c)
			}
		}
	}
}

func TestCardListPutWithExistingCard(t *testing.T) {
	cards := CardList{}
	expected := CardList{}

	fill(&cards)
	fill(&expected)

	if !reflect.DeepEqual(cards, expected) {
		t.Errorf("Invalid CardList\n   expected: %v\n   got:      %v", expected, cards)
	}

	card := Card{
		CardNumber: uint32(6000005),
		From:       date("2020-02-03"),
		To:         date("2020-11-30"),
		Doors: map[uint8]uint8{
			1: 0,
			2: 1,
			3: 1,
			4: 0,
		},
	}

	expected[5] = &card

	err := cards.Put(&card)
	if err != nil {
		t.Fatalf("Unexpected error adding card to list: %v", err)
	}

	if !reflect.DeepEqual(cards, expected) {
		for i, c := range cards {
			if !reflect.DeepEqual(c, expected[i]) {
				t.Errorf("Invalid CardList entry %v\n   expected: %v\n   got:      %v", i, *expected[i], *c)
			}
		}
	}
}

func TestCardListPutWithFullList(t *testing.T) {
	cards := CardList{}

	for i := 0; i < len(cards); i++ {
		cards[i] = &Card{
			CardNumber: uint32(6000000 + i),
			From:       date("2020-01-01"),
			To:         date("2020-12-31"),
			Doors: map[uint8]uint8{
				1: 1,
				2: 0,
				3: 0,
				4: 1,
			},
		}
	}

	card := Card{
		CardNumber: uint32(5000019),
		From:       date("2020-01-01"),
		To:         date("2020-12-31"),
		Doors: map[uint8]uint8{
			1: 0,
			2: 1,
			3: 1,
			4: 0,
		},
	}

	err := cards.Put(&card)
	if err == nil {
		t.Fatalf("Expected error adding card to full list, got %v", err)
	}
}

func TestCardListDeleteActiveCard(t *testing.T) {
	cards := CardList{}
	expected := CardList{}

	fill(&cards)
	fill(&expected)

	expected[5] = &Card{
		CardNumber: 0xffffffff,
	}

	ok := cards.Delete(6000005)
	if ok != true {
		t.Fatalf("Unexpected result deleting card - expected:%v, got %v", true, ok)
	}

	if !reflect.DeepEqual(cards, expected) {
		for i, c := range cards {
			if !reflect.DeepEqual(c, expected[i]) {
				t.Errorf("Invalid CardList entry %v\n   expected: %v\n   got:      %v", i, *expected[i], *c)
			}
		}
	}
}

func TestCardListDeleteCardNotInList(t *testing.T) {
	cards := CardList{}
	expected := CardList{}

	fill(&cards)
	fill(&expected)

	ok := cards.Delete(5000019)
	if ok != false {
		t.Fatalf("Unexpected result deleting card - expected:%v, got %v", false, ok)
	}

	if !reflect.DeepEqual(cards, expected) {
		for i, c := range cards {
			if !reflect.DeepEqual(c, expected[i]) {
				t.Errorf("Invalid CardList entry %v\n   expected: %v\n   got:      %v", i, *expected[i], *c)
			}
		}
	}
}

func TestCardListDeleteAll(t *testing.T) {
	cards := CardList{}

	ok := cards.DeleteAll()
	if ok != true {
		t.Fatalf("Unexpected result deleting card list - expected:%v, got %v", true, ok)
	}

	if N := cards.Size(); N != 0 {
		t.Errorf("CardList.Size returned %v, expected: %v", N, 0)
	}

	for i, c := range cards {
		if c != nil {
			t.Errorf("Invalid CardList entry %v\n   expected: %v\n   got:      %v", i, nil, *c)
		}
	}
}

func fill(l *CardList) uint32 {
	for i := 0; i < 29; i++ {
		l[i] = &Card{
			CardNumber: uint32(6000000 + i),
			From:       date("2020-01-01"),
			To:         date("2020-12-31"),
			Doors: map[uint8]uint8{
				1: 1,
				2: 0,
				3: 0,
				4: 1,
			},
		}
	}

	l[11] = nil
	l[17] = nil

	l[19].CardNumber = 0

	l[3].CardNumber = 0xffffffff
	l[15].CardNumber = 0xffffffff
	l[25].CardNumber = 0xffffffff

	return 23
}
