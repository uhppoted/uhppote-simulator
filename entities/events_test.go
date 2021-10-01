package entities

import (
	"testing"
)

func TestSetIndex(t *testing.T) {
	events := EventList{

		Size:  8,
		First: 1,
		Last:  5,
		Index: 3,
		Events: []Event{
			Event{Index: 1},
			Event{Index: 2},
			Event{Index: 3},
		},
	}

	if !events.SetIndex(4) {
		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", true, false)
	}

	if events.Index != 4 {
		t.Errorf("SetIndex failed to update internal index - expected:%v, got:%v", 4, events.Index)
	}
}

func TestSetIndexWithOutOfRangeValue(t *testing.T) {
	events := EventList{

		Size:  8,
		First: 1,
		Last:  5,
		Index: 3,
		Events: []Event{
			Event{Index: 1},
			Event{Index: 2},
			Event{Index: 3},
		},
	}

	if events.SetIndex(6) {
		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", false, true)
	}

	if events.Index != 3 {
		t.Errorf("SetIndex incorrected updated internal index - expected:%v, got:%v", 3, events.Index)
	}
}

// FIXME (provisional - pending validation against controller)
func TestSetIndexWithRollover(t *testing.T) {
	events := EventList{

		Size:   32,
		First:  27,
		Last:   5,
		Index:  20,
		Events: []Event{},
	}

	if !events.SetIndex(34) {
		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", true, true)
	}

	if events.Index != 1 {
		t.Errorf("SetIndex incorrected updated internal index - expected:%v, got:%v", 1, events.Index)
	}
}

func TestSetIndexWithRolloverAndOutOfRange(t *testing.T) {
	events := EventList{

		Size:   32,
		First:  27,
		Last:   5,
		Index:  20,
		Events: []Event{},
	}

	if events.SetIndex(6) {
		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", false, true)
	}

	if events.Index != 20 {
		t.Errorf("SetIndex incorrected updated internal index - expected:%v, got:%v", 20, events.Index)
	}
}
