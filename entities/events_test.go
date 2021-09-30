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
