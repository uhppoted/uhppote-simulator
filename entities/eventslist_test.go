package entities

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/uhppoted/uhppote-core/types"
)

func TestNewEventList(t *testing.T) {
	expected := EventList{
		size:   256,
		chunk:  8,
		index:  0,
		events: []Event{},
	}

	events := NewEventList()

	if !reflect.DeepEqual(events, expected) {
		t.Errorf("incorrect EventList\n   expected:%+v\n   got:    %+v", expected, events)
	}
}

func TestMarshalEventList(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	l := EventList{
		size:  64,
		chunk: 8,
		index: 19,
		events: []Event{
			Event{Index: 1, Timestamp: &timestamp},
			Event{Index: 2, Timestamp: &timestamp},
			Event{Index: 3, Timestamp: &timestamp},
		},
	}

	expected := `{
  "size": 64,
  "chunk": 8,
  "index": 19,
  "events": [
    {
      "index": 1,
      "type": 0,
      "granted": false,
      "door": 0,
      "direction": 0,
      "card": 0,
      "timestamp": "2021-12-27 13:14:15",
      "reason": 0
    },
    {
      "index": 2,
      "type": 0,
      "granted": false,
      "door": 0,
      "direction": 0,
      "card": 0,
      "timestamp": "2021-12-27 13:14:15",
      "reason": 0
    },
    {
      "index": 3,
      "type": 0,
      "granted": false,
      "door": 0,
      "direction": 0,
      "card": 0,
      "timestamp": "2021-12-27 13:14:15",
      "reason": 0
    }
  ]
}`

	b, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		t.Fatalf("Unexpected error marshalling EventList (%v)", err)
	}

	if string(b) != expected {
		t.Errorf("EventList incorrectly marshalled\n   expected:%v\n   got:     %v", expected, string(b))
	}
}

func TestUnmarshalEventList(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	bytes := []byte(`{
  "size": 64,
  "chunk": 8,
  "index": 19,
  "events": [
    { "index": 1, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 2, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 3, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 }
  ]
}`)

	expected := EventList{
		size:  64,
		chunk: 8,
		index: 19,
		events: []Event{
			Event{Index: 1, Timestamp: &timestamp},
			Event{Index: 2, Timestamp: &timestamp},
			Event{Index: 3, Timestamp: &timestamp},
		},
	}

	l := EventList{}
	if err := json.Unmarshal(bytes, &l); err != nil {
		t.Fatalf("Unexpected error unmarshalling EventList (%v)", err)
	}

	if !reflect.DeepEqual(l, expected) {
		t.Errorf("EventList incorrectly unmarshalled\n   expected:%#v\n   got:     %#v", expected, l)
	}
}

func TestUnmarshalEventListWithUnOrderedEvents(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	bytes := []byte(`{
  "size": 64,
  "chunk": 8,
  "index": 19,
  "events": [
    { "index": 1, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 3, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 2, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 }
  ]
}`)

	expected := EventList{
		size:  64,
		chunk: 8,
		index: 19,
		events: []Event{
			Event{Index: 1, Timestamp: &timestamp},
			Event{Index: 2, Timestamp: &timestamp},
			Event{Index: 3, Timestamp: &timestamp},
		},
	}

	l := EventList{}
	if err := json.Unmarshal(bytes, &l); err != nil {
		t.Fatalf("Unexpected error unmarshalling EventList (%v)", err)
	}

	if !reflect.DeepEqual(l, expected) {
		t.Errorf("EventList incorrectly unmarshalled\n   expected:%#v\n   got:     %#v", expected, l)
	}
}

func TestUnmarshalEventListWithTooManyEvents(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	bytes := []byte(`{
  "size": 8,
  "chunk": 2,
  "index": 19,
  "events": [
    { "index": 1,  "type": 101, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 2,  "type": 102, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 3,  "type": 103, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 4,  "type": 104, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 5,  "type": 105, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 6,  "type": 106, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 7,  "type": 107, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 8,  "type": 108, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 9,  "type": 109, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 10, "type": 110, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 11, "type": 111, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 12, "type": 112, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 13, "type": 113, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 }
  ]
}`)

	expected := EventList{
		size:  8,
		chunk: 2,
		index: 19,
		events: []Event{
			Event{Index: 7, Type: 107, Timestamp: &timestamp},
			Event{Index: 8, Type: 108, Timestamp: &timestamp},
			Event{Index: 9, Type: 109, Timestamp: &timestamp},
			Event{Index: 10, Type: 110, Timestamp: &timestamp},
			Event{Index: 11, Type: 111, Timestamp: &timestamp},
			Event{Index: 12, Type: 112, Timestamp: &timestamp},
			Event{Index: 13, Type: 113, Timestamp: &timestamp},
		},
	}

	l := EventList{}
	if err := json.Unmarshal(bytes, &l); err != nil {
		t.Fatalf("Unexpected error unmarshalling EventList (%v)", err)
	}

	if !reflect.DeepEqual(l, expected) {
		t.Errorf("EventList incorrectly unmarshalled\n   expected:%#v\n   got:     %#v", expected, l)
	}
}

func TestUnmarshalEventListWithDefaultSizeAndChunk(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	bytes := []byte(`{
  "index": 19,
  "events": [
    { "index": 1, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 2, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 3, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 }
  ]
}`)

	expected := EventList{
		size:  256,
		chunk: 8,
		index: 19,
		events: []Event{
			Event{Index: 1, Timestamp: &timestamp},
			Event{Index: 2, Timestamp: &timestamp},
			Event{Index: 3, Timestamp: &timestamp},
		},
	}

	l := EventList{}
	if err := json.Unmarshal(bytes, &l); err != nil {
		t.Fatalf("Unexpected error unmarshalling EventList (%v)", err)
	}

	if !reflect.DeepEqual(l, expected) {
		t.Errorf("EventList incorrectly unmarshalled\n   expected:%#v\n   got:     %#v", expected, l)
	}
}

func TestUnmarshalEventListWithZeroSizeAndChunk(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	bytes := []byte(`{
  "size": 0,
  "chunk": 0,
  "index": 19,
  "events": [
    { "index": 1, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 2, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 3, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 }
  ]
}`)

	expected := EventList{
		size:  256,
		chunk: 8,
		index: 19,
		events: []Event{
			Event{Index: 1, Timestamp: &timestamp},
			Event{Index: 2, Timestamp: &timestamp},
			Event{Index: 3, Timestamp: &timestamp},
		},
	}

	l := EventList{}
	if err := json.Unmarshal(bytes, &l); err != nil {
		t.Fatalf("Unexpected error unmarshalling EventList (%v)", err)
	}

	if !reflect.DeepEqual(l, expected) {
		t.Errorf("EventList incorrectly unmarshalled\n   expected:%#v\n   got:     %#v", expected, l)
	}
}

func TestAddEvent(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	expected := EventList{
		size:  64,
		chunk: 8,
		index: 19,
		events: []Event{
			Event{Index: 1, Timestamp: &timestamp, Type: 101},
			Event{Index: 2, Timestamp: &timestamp, Type: 102},
			Event{Index: 3, Timestamp: &timestamp, Type: 103},
			Event{Index: 4, Timestamp: &timestamp, Type: 104},
		},
	}

	events := EventList{
		size:  64,
		chunk: 8,
		index: 19,
		events: []Event{
			Event{Index: 1, Timestamp: &timestamp, Type: 101},
			Event{Index: 2, Timestamp: &timestamp, Type: 102},
			Event{Index: 3, Timestamp: &timestamp, Type: 103},
		},
	}

	event := Event{Timestamp: &timestamp, Type: 104}

	index := events.Add(event)

	if index != 4 {
		t.Errorf("Incorrect EventList index from Add - expected:%v, got:%v", 4, index)
	}

	if !reflect.DeepEqual(events, expected) {
		t.Errorf("Incorrect EventList after Add\n   expected:%#v\n   got:     %#v", expected, events)
	}
}

func TestAddEventWithEmptyList(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	expected := EventList{
		size:  64,
		chunk: 8,
		index: 19,
		events: []Event{
			Event{Index: 1, Timestamp: &timestamp, Type: 105},
		},
	}

	events := EventList{
		size:   64,
		chunk:  8,
		index:  19,
		events: []Event{},
	}

	event := Event{Timestamp: &timestamp, Type: 105}

	index := events.Add(event)

	if index != 1 {
		t.Errorf("Incorrect EventList index from Add - expected:%v, got:%v", 1, index)
	}

	if !reflect.DeepEqual(events, expected) {
		t.Errorf("Incorrect EventList after Add\n   expected:%#v\n   got:     %#v", expected, events)
	}
}

func TestAddEventWithFullList(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	expected := EventList{
		size:  8,
		chunk: 2,
		index: 19,
		events: []Event{
			Event{Index: 3, Timestamp: &timestamp, Type: 103},
			Event{Index: 4, Timestamp: &timestamp, Type: 104},
			Event{Index: 5, Timestamp: &timestamp, Type: 105},
			Event{Index: 6, Timestamp: &timestamp, Type: 106},
			Event{Index: 7, Timestamp: &timestamp, Type: 107},
			Event{Index: 8, Timestamp: &timestamp, Type: 108},
			Event{Index: 9, Timestamp: &timestamp, Type: 109},
		},
	}

	events := EventList{
		size:  8,
		chunk: 2,
		index: 19,
		events: []Event{
			Event{Index: 1, Timestamp: &timestamp, Type: 101},
			Event{Index: 2, Timestamp: &timestamp, Type: 102},
			Event{Index: 3, Timestamp: &timestamp, Type: 103},
			Event{Index: 4, Timestamp: &timestamp, Type: 104},
			Event{Index: 5, Timestamp: &timestamp, Type: 105},
			Event{Index: 6, Timestamp: &timestamp, Type: 106},
			Event{Index: 7, Timestamp: &timestamp, Type: 107},
			Event{Index: 8, Timestamp: &timestamp, Type: 108},
		},
	}

	event := Event{Timestamp: &timestamp, Type: 109}

	index := events.Add(event)

	if index != 9 {
		t.Errorf("Incorrect EventList index from Add - expected:%v, got:%v", 9, index)
	}

	if !reflect.DeepEqual(events, expected) {
		t.Errorf("Incorrect EventList after Add\n   expected:%#v\n   got:     %#v", expected, events)
	}
}

func TestGetEventWithNoEvents(t *testing.T) {
	events := EventList{
		size:   8,
		index:  0,
		events: []Event{},
	}

	expected := Event{}

	e := events.Get(123)

	if !reflect.DeepEqual(e, expected) {
		t.Errorf("Incorrect 'no event' return for empty EventList\n   expected:%#v\n   got:     %#v", expected, e)
	}
}

func TestGetFirstEvent(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	events := EventList{
		size:  8,
		index: 0,
		events: []Event{
			Event{Index: 1001, Timestamp: &timestamp, Type: 11},
			Event{Index: 1002, Timestamp: &timestamp, Type: 12},
			Event{Index: 1003, Timestamp: &timestamp, Type: 13},
			Event{Index: 1004, Timestamp: &timestamp, Type: 14},
			Event{Index: 1005, Timestamp: &timestamp, Type: 15},
			Event{Index: 1006, Timestamp: &timestamp, Type: 16},
			Event{Index: 1007, Timestamp: &timestamp, Type: 17},
			Event{Index: 1008, Timestamp: &timestamp, Type: 18},
		},
	}

	expected := Event{Index: 1001, Timestamp: &timestamp, Type: 11}

	e := events.Get(0)

	if !reflect.DeepEqual(e, expected) {
		t.Errorf("Incorrect 'first event'\n   expected:%#v\n   got:     %#v", expected, e)
	}
}

func TestGetLastEvent(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	events := EventList{
		size:  8,
		index: 0,
		events: []Event{
			Event{Index: 1001, Timestamp: &timestamp, Type: 11},
			Event{Index: 1002, Timestamp: &timestamp, Type: 12},
			Event{Index: 1003, Timestamp: &timestamp, Type: 13},
			Event{Index: 1004, Timestamp: &timestamp, Type: 14},
			Event{Index: 1005, Timestamp: &timestamp, Type: 15},
			Event{Index: 1006, Timestamp: &timestamp, Type: 16},
			Event{Index: 1007, Timestamp: &timestamp, Type: 17},
			Event{Index: 1008, Timestamp: &timestamp, Type: 18},
		},
	}

	expected := Event{Index: 1008, Timestamp: &timestamp, Type: 18}

	e := events.Get(0xffffffff)

	if !reflect.DeepEqual(e, expected) {
		t.Errorf("Incorrect 'last event'\n   expected:%#v\n   got:     %#v", expected, e)
	}
}

func TestGetEventAtIndex(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	events := EventList{
		size:  8,
		index: 0,
		events: []Event{
			Event{Index: 1001, Timestamp: &timestamp, Type: 11},
			Event{Index: 1002, Timestamp: &timestamp, Type: 12},
			Event{Index: 1003, Timestamp: &timestamp, Type: 13},
			Event{Index: 1004, Timestamp: &timestamp, Type: 14},
			Event{Index: 1005, Timestamp: &timestamp, Type: 15},
			Event{Index: 1006, Timestamp: &timestamp, Type: 16},
			Event{Index: 1007, Timestamp: &timestamp, Type: 17},
			Event{Index: 1008, Timestamp: &timestamp, Type: 18},
		},
	}

	expected := Event{Index: 1003, Timestamp: &timestamp, Type: 13}

	e := events.Get(1003)

	if !reflect.DeepEqual(e, expected) {
		t.Errorf("Incorrect event\n   expected:%#v\n   got:     %#v", expected, e)
	}
}

func TestGetOverwrittenEvent(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	events := EventList{
		size:  8,
		index: 0,
		events: []Event{
			Event{Index: 1001, Timestamp: &timestamp, Type: 11},
			Event{Index: 1002, Timestamp: &timestamp, Type: 12},
			Event{Index: 1003, Timestamp: &timestamp, Type: 13},
			Event{Index: 1004, Timestamp: &timestamp, Type: 14},
			Event{Index: 1005, Timestamp: &timestamp, Type: 15},
			Event{Index: 1006, Timestamp: &timestamp, Type: 16},
			Event{Index: 1007, Timestamp: &timestamp, Type: 17},
			Event{Index: 1008, Timestamp: &timestamp, Type: 18},
		},
	}

	expected := Event{Type: 0xff}

	e := events.Get(117)

	if !reflect.DeepEqual(e, expected) {
		t.Errorf("Incorrect 'overwritten event'\n   expected:%#v\n   got:     %#v", expected, e)
	}
}

func TestGetOutOfRangeEvent(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	events := EventList{
		size:  8,
		index: 0,
		events: []Event{
			Event{Index: 1001, Timestamp: &timestamp, Type: 11},
			Event{Index: 1002, Timestamp: &timestamp, Type: 12},
			Event{Index: 1003, Timestamp: &timestamp, Type: 13},
			Event{Index: 1004, Timestamp: &timestamp, Type: 14},
			Event{Index: 1005, Timestamp: &timestamp, Type: 15},
			Event{Index: 1006, Timestamp: &timestamp, Type: 16},
			Event{Index: 1007, Timestamp: &timestamp, Type: 17},
			Event{Index: 1008, Timestamp: &timestamp, Type: 18},
		},
	}

	expected := Event{}

	e := events.Get(12345)

	if !reflect.DeepEqual(e, expected) {
		t.Errorf("Incorrect 'out of range event'\n   expected:%#v\n   got:     %#v", expected, e)
	}
}

func TestSetIndex(t *testing.T) {
	events := EventList{
		size:  8,
		index: 3,
		events: []Event{
			Event{Index: 1},
			Event{Index: 2},
			Event{Index: 3},
			Event{Index: 4},
		},
	}

	if !events.SetIndex(4) {
		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", true, false)
	}

	if events.index != 4 {
		t.Errorf("SetIndex failed to update internal index - expected:%v, got:%v", 4, events.index)
	}
}

func TestSetIndexWithSameValue(t *testing.T) {
	events := EventList{
		size:  8,
		index: 3,
		events: []Event{
			Event{Index: 1},
			Event{Index: 2},
			Event{Index: 3},
			Event{Index: 4},
		},
	}

	if events.SetIndex(3) {
		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", false, true)
	}

	if events.index != 3 {
		t.Errorf("SetIndex updated internal index - expected:%v, got:%v", 3, events.index)
	}
}

func TestSetIndexWithZero(t *testing.T) {
	events := EventList{
		size:  8,
		index: 3,
		events: []Event{
			Event{Index: 1},
			Event{Index: 2},
			Event{Index: 3},
		},
	}

	if !events.SetIndex(0) {
		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", true, false)
	}

	if events.index != 0 {
		t.Errorf("SetIndex failed to update internal index - expected:%v, got:%v", 0, events.index)
	}
}

func TestSetIndexWithAlreadyZero(t *testing.T) {
	events := EventList{
		size:  8,
		index: 0,
		events: []Event{
			Event{Index: 1},
			Event{Index: 2},
			Event{Index: 3},
		},
	}

	if events.SetIndex(0) {
		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", false, true)
	}

	if events.index != 0 {
		t.Errorf("SetIndex updated internal index - expected:%v, got:%v", 0, events.index)
	}
}

func TestSetIndexWithLessThanFirstIndex(t *testing.T) {
	events := EventList{
		size:  8,
		index: 1002,
		events: []Event{
			Event{Index: 1001},
			Event{Index: 1002},
			Event{Index: 1003},
		},
	}

	if !events.SetIndex(123) {
		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", false, true)
	}

	if events.index != 123 {
		t.Errorf("SetIndex incorrected updated internal index - expected:%v, got:%v", 123, events.index)
	}
}

func TestSetIndexWithLastIndex(t *testing.T) {
	events := EventList{
		size:  8,
		index: 123,
		events: []Event{
			Event{Index: 1001},
			Event{Index: 1002},
			Event{Index: 1003},
		},
	}

	if !events.SetIndex(1003) {
		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", false, true)
	}

	if events.index != 1003 {
		t.Errorf("SetIndex incorrected updated internal index - expected:%v, got:%v", 1003, events.index)
	}
}

func TestSetIndexWithGreaterThanLastIndex(t *testing.T) {
	events := EventList{
		size:  8,
		index: 123,
		events: []Event{
			Event{Index: 1001},
			Event{Index: 1002},
			Event{Index: 1003},
		},
	}

	if events.SetIndex(1006) {
		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", false, true)
	}

	if events.index != 123 {
		t.Errorf("SetIndex incorrected updated internal index - expected:%v, got:%v", 123, events.index)
	}
}
