package entities

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/uhppoted/uhppote-core/types"
)

func TestMarshalEvent(t *testing.T) {
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 9, time.Local)

	e := Event{
		Index:     7,
		Type:      13,
		Granted:   true,
		Door:      3,
		Direction: 1,
		Card:      8165535,
		Timestamp: types.DateTime(timestamp),
		Reason:    6,
	}

	expected := `{
  "index": 7,
  "type": 13,
  "granted": true,
  "door": 3,
  "direction": 1,
  "card": 8165535,
  "timestamp": "2021-12-27 13:14:15",
  "reason": 6
}`

	b, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		t.Fatalf("Unexpected error marshalling Event (%v)", err)
	}

	if string(b) != expected {
		t.Errorf("Event incorrectly marshalled\n   expected:%v\n   got:     %v", expected, string(b))
	}
}

func TestUnmarshalEvent(t *testing.T) {
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

	expected := Event{
		Index:     7,
		Type:      13,
		Granted:   true,
		Door:      3,
		Direction: 1,
		Card:      8165535,
		Timestamp: types.DateTime(timestamp),
		Reason:    6,
	}

	bytes := []byte(`{ "index":7, "type":13, "granted":true, "door":3, "direction":1, "card":8165535, "timestamp":"2021-12-27 13:14:15", "reason":6 }`)

	var event Event
	if err := json.Unmarshal(bytes, &event); err != nil {
		t.Fatalf("Unexpected error unmarshalling Event (%v)", err)
	}

	if !reflect.DeepEqual(event, expected) {
		t.Errorf("Event incorrectly unmarshalled\n   expected:%#v\n   got:     %#v", expected, event)
	}
}

func TestMarshalEventList(t *testing.T) {
	l := EventList{
		Size:  64,
		Chunk: 8,
		Index: 19,
		Events: []Event{
			Event{Index: 1},
			Event{Index: 2},
			Event{Index: 3},
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
      "timestamp": "0001-01-01 00:00:00",
      "reason": 0
    },
    {
      "index": 2,
      "type": 0,
      "granted": false,
      "door": 0,
      "direction": 0,
      "card": 0,
      "timestamp": "0001-01-01 00:00:00",
      "reason": 0
    },
    {
      "index": 3,
      "type": 0,
      "granted": false,
      "door": 0,
      "direction": 0,
      "card": 0,
      "timestamp": "0001-01-01 00:00:00",
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
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

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
		Size:  64,
		Chunk: 8,
		Index: 19,
		Events: []Event{
			Event{Index: 1, Timestamp: types.DateTime(timestamp)},
			Event{Index: 2, Timestamp: types.DateTime(timestamp)},
			Event{Index: 3, Timestamp: types.DateTime(timestamp)},
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
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

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
		Size:  64,
		Chunk: 8,
		Index: 19,
		Events: []Event{
			Event{Index: 1, Timestamp: types.DateTime(timestamp)},
			Event{Index: 2, Timestamp: types.DateTime(timestamp)},
			Event{Index: 3, Timestamp: types.DateTime(timestamp)},
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
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

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
    { "index": 12, "type": 112, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 }
  ]
}`)

	expected := EventList{
		Size:  8,
		Chunk: 2,
		Index: 19,
		Events: []Event{
			Event{Index: 5, Type: 105, Timestamp: types.DateTime(timestamp)},
			Event{Index: 6, Type: 106, Timestamp: types.DateTime(timestamp)},
			Event{Index: 7, Type: 107, Timestamp: types.DateTime(timestamp)},
			Event{Index: 8, Type: 108, Timestamp: types.DateTime(timestamp)},
			Event{Index: 9, Type: 109, Timestamp: types.DateTime(timestamp)},
			Event{Index: 10, Type: 110, Timestamp: types.DateTime(timestamp)},
			Event{Index: 11, Type: 111, Timestamp: types.DateTime(timestamp)},
			Event{Index: 12, Type: 112, Timestamp: types.DateTime(timestamp)},
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
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

	bytes := []byte(`{
  "index": 19,
  "events": [
    { "index": 1, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 2, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 },
    { "index": 3, "type": 0, "granted": false, "door": 0, "direction": 0, "card": 0, "timestamp": "2021-12-27 13:14:15", "reason": 0 }
  ]
}`)

	expected := EventList{
		Size:  256,
		Chunk: 8,
		Index: 19,
		Events: []Event{
			Event{Index: 1, Timestamp: types.DateTime(timestamp)},
			Event{Index: 2, Timestamp: types.DateTime(timestamp)},
			Event{Index: 3, Timestamp: types.DateTime(timestamp)},
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
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

	expected := EventList{
		Size:  64,
		Chunk: 8,
		Index: 19,
		Events: []Event{
			Event{Index: 1, Timestamp: types.DateTime(timestamp), Type: 101},
			Event{Index: 2, Timestamp: types.DateTime(timestamp), Type: 102},
			Event{Index: 3, Timestamp: types.DateTime(timestamp), Type: 103},
			Event{Index: 4, Timestamp: types.DateTime(timestamp), Type: 104},
		},
	}

	events := EventList{
		Size:  64,
		Chunk: 8,
		Index: 19,
		Events: []Event{
			Event{Index: 1, Timestamp: types.DateTime(timestamp), Type: 101},
			Event{Index: 2, Timestamp: types.DateTime(timestamp), Type: 102},
			Event{Index: 3, Timestamp: types.DateTime(timestamp), Type: 103},
		},
	}

	event := Event{Timestamp: types.DateTime(timestamp), Type: 104}

	index := events.Add(event)

	if index != 4 {
		t.Errorf("Incorrect EventList index from Add - expected:%v, got:%v", 4, index)
	}

	if !reflect.DeepEqual(events, expected) {
		t.Errorf("Incorrect EventList after Add\n   expected:%#v\n   got:     %#v", expected, events)
	}
}

func TestAddEventWithEmptyList(t *testing.T) {
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

	expected := EventList{
		Size:  64,
		Chunk: 8,
		Index: 19,
		Events: []Event{
			Event{Index: 1, Timestamp: types.DateTime(timestamp), Type: 105},
		},
	}

	events := EventList{
		Size:   64,
		Chunk:  8,
		Index:  19,
		Events: []Event{},
	}

	event := Event{Timestamp: types.DateTime(timestamp), Type: 105}

	index := events.Add(event)

	if index != 1 {
		t.Errorf("Incorrect EventList index from Add - expected:%v, got:%v", 1, index)
	}

	if !reflect.DeepEqual(events, expected) {
		t.Errorf("Incorrect EventList after Add\n   expected:%#v\n   got:     %#v", expected, events)
	}
}

func TestAddEventWithFullList(t *testing.T) {
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

	expected := EventList{
		Size:  8,
		Chunk: 2,
		Index: 19,
		Events: []Event{
			Event{Index: 3, Timestamp: types.DateTime(timestamp), Type: 103},
			Event{Index: 4, Timestamp: types.DateTime(timestamp), Type: 104},
			Event{Index: 5, Timestamp: types.DateTime(timestamp), Type: 105},
			Event{Index: 6, Timestamp: types.DateTime(timestamp), Type: 106},
			Event{Index: 7, Timestamp: types.DateTime(timestamp), Type: 107},
			Event{Index: 8, Timestamp: types.DateTime(timestamp), Type: 108},
			Event{Index: 9, Timestamp: types.DateTime(timestamp), Type: 109},
		},
	}

	events := EventList{
		Size:  8,
		Chunk: 2,
		Index: 19,
		Events: []Event{
			Event{Index: 1, Timestamp: types.DateTime(timestamp), Type: 101},
			Event{Index: 2, Timestamp: types.DateTime(timestamp), Type: 102},
			Event{Index: 3, Timestamp: types.DateTime(timestamp), Type: 103},
			Event{Index: 4, Timestamp: types.DateTime(timestamp), Type: 104},
			Event{Index: 5, Timestamp: types.DateTime(timestamp), Type: 105},
			Event{Index: 6, Timestamp: types.DateTime(timestamp), Type: 106},
			Event{Index: 7, Timestamp: types.DateTime(timestamp), Type: 107},
			Event{Index: 8, Timestamp: types.DateTime(timestamp), Type: 108},
		},
	}

	event := Event{Timestamp: types.DateTime(timestamp), Type: 109}

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
		Size:   8,
		Index:  0,
		Events: []Event{},
	}

	expected := Event{}

	e := events.Get(123)

	if !reflect.DeepEqual(e, expected) {
		t.Errorf("Incorrect 'no event' return for empty EventList\n   expected:%#v\n   got:     %#v", expected, e)
	}
}

func TestGetFirstEvent(t *testing.T) {
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

	events := EventList{
		Size:  8,
		Index: 0,
		Events: []Event{
			Event{Index: 1001, Timestamp: types.DateTime(timestamp), Type: 11},
			Event{Index: 1002, Timestamp: types.DateTime(timestamp), Type: 12},
			Event{Index: 1003, Timestamp: types.DateTime(timestamp), Type: 13},
			Event{Index: 1004, Timestamp: types.DateTime(timestamp), Type: 14},
			Event{Index: 1005, Timestamp: types.DateTime(timestamp), Type: 15},
			Event{Index: 1006, Timestamp: types.DateTime(timestamp), Type: 16},
			Event{Index: 1007, Timestamp: types.DateTime(timestamp), Type: 17},
			Event{Index: 1008, Timestamp: types.DateTime(timestamp), Type: 18},
		},
	}

	expected := Event{Index: 1001, Timestamp: types.DateTime(timestamp), Type: 11}

	e := events.Get(0)

	if !reflect.DeepEqual(e, expected) {
		t.Errorf("Incorrect 'first event'\n   expected:%#v\n   got:     %#v", expected, e)
	}
}

func TestGetLastEvent(t *testing.T) {
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

	events := EventList{
		Size:  8,
		Index: 0,
		Events: []Event{
			Event{Index: 1001, Timestamp: types.DateTime(timestamp), Type: 11},
			Event{Index: 1002, Timestamp: types.DateTime(timestamp), Type: 12},
			Event{Index: 1003, Timestamp: types.DateTime(timestamp), Type: 13},
			Event{Index: 1004, Timestamp: types.DateTime(timestamp), Type: 14},
			Event{Index: 1005, Timestamp: types.DateTime(timestamp), Type: 15},
			Event{Index: 1006, Timestamp: types.DateTime(timestamp), Type: 16},
			Event{Index: 1007, Timestamp: types.DateTime(timestamp), Type: 17},
			Event{Index: 1008, Timestamp: types.DateTime(timestamp), Type: 18},
		},
	}

	expected := Event{Index: 1008, Timestamp: types.DateTime(timestamp), Type: 18}

	e := events.Get(0xffffffff)

	if !reflect.DeepEqual(e, expected) {
		t.Errorf("Incorrect 'last event'\n   expected:%#v\n   got:     %#v", expected, e)
	}
}

func TestGetEventAtIndex(t *testing.T) {
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

	events := EventList{
		Size:  8,
		Index: 0,
		Events: []Event{
			Event{Index: 1001, Timestamp: types.DateTime(timestamp), Type: 11},
			Event{Index: 1002, Timestamp: types.DateTime(timestamp), Type: 12},
			Event{Index: 1003, Timestamp: types.DateTime(timestamp), Type: 13},
			Event{Index: 1004, Timestamp: types.DateTime(timestamp), Type: 14},
			Event{Index: 1005, Timestamp: types.DateTime(timestamp), Type: 15},
			Event{Index: 1006, Timestamp: types.DateTime(timestamp), Type: 16},
			Event{Index: 1007, Timestamp: types.DateTime(timestamp), Type: 17},
			Event{Index: 1008, Timestamp: types.DateTime(timestamp), Type: 18},
		},
	}

	expected := Event{Index: 1003, Timestamp: types.DateTime(timestamp), Type: 13}

	e := events.Get(1003)

	if !reflect.DeepEqual(e, expected) {
		t.Errorf("Incorrect event\n   expected:%#v\n   got:     %#v", expected, e)
	}
}

func TestGetOverwrittenEvent(t *testing.T) {
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

	events := EventList{
		Size:  8,
		Index: 0,
		Events: []Event{
			Event{Index: 1001, Timestamp: types.DateTime(timestamp), Type: 11},
			Event{Index: 1002, Timestamp: types.DateTime(timestamp), Type: 12},
			Event{Index: 1003, Timestamp: types.DateTime(timestamp), Type: 13},
			Event{Index: 1004, Timestamp: types.DateTime(timestamp), Type: 14},
			Event{Index: 1005, Timestamp: types.DateTime(timestamp), Type: 15},
			Event{Index: 1006, Timestamp: types.DateTime(timestamp), Type: 16},
			Event{Index: 1007, Timestamp: types.DateTime(timestamp), Type: 17},
			Event{Index: 1008, Timestamp: types.DateTime(timestamp), Type: 18},
		},
	}

	expected := Event{Type: 0xff}

	e := events.Get(117)

	if !reflect.DeepEqual(e, expected) {
		t.Errorf("Incorrect 'overwritten event'\n   expected:%#v\n   got:     %#v", expected, e)
	}
}

func TestGetOutOfRangeEvent(t *testing.T) {
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

	events := EventList{
		Size:  8,
		Index: 0,
		Events: []Event{
			Event{Index: 1001, Timestamp: types.DateTime(timestamp), Type: 11},
			Event{Index: 1002, Timestamp: types.DateTime(timestamp), Type: 12},
			Event{Index: 1003, Timestamp: types.DateTime(timestamp), Type: 13},
			Event{Index: 1004, Timestamp: types.DateTime(timestamp), Type: 14},
			Event{Index: 1005, Timestamp: types.DateTime(timestamp), Type: 15},
			Event{Index: 1006, Timestamp: types.DateTime(timestamp), Type: 16},
			Event{Index: 1007, Timestamp: types.DateTime(timestamp), Type: 17},
			Event{Index: 1008, Timestamp: types.DateTime(timestamp), Type: 18},
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
		Size:  8,
		Index: 3,
		Events: []Event{
			Event{Index: 1},
			Event{Index: 2},
			Event{Index: 3},
			Event{Index: 4},
		},
	}

	if !events.SetIndex(4) {
		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", true, false)
	}

	if events.Index != 4 {
		t.Errorf("SetIndex failed to update internal index - expected:%v, got:%v", 4, events.Index)
	}
}

func TestSetIndexWithZero(t *testing.T) {
	events := EventList{
		Size:  8,
		Index: 3,
		Events: []Event{
			Event{Index: 1},
			Event{Index: 2},
			Event{Index: 3},
		},
	}

	if !events.SetIndex(0) {
		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", true, false)
	}

	if events.Index != 0 {
		t.Errorf("SetIndex failed to update internal index - expected:%v, got:%v", 0, events.Index)
	}
}

func TestSetIndexWithOutOfRangeValue(t *testing.T) {
	events := EventList{
		Size:  8,
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

// // FIXME (provisional - pending validation against controller)
// func TestSetIndexWithRollover(t *testing.T) {
// 	events := EventList{
//
// 		Size:   32,
// 		First:  27,
// 		Last:   5,
// 		Index:  20,
// 		Events: []Event{},
// 	}
//
// 	if !events.SetIndex(34) {
// 		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", true, true)
// 	}
//
// 	if events.Index != 1 {
// 		t.Errorf("SetIndex incorrected updated internal index - expected:%v, got:%v", 1, events.Index)
// 	}
// }

// func TestSetIndexWithRolloverAndOutOfRange(t *testing.T) {
// 	events := EventList{
//
// 		Size:   32,
// 		First:  27,
// 		Last:   5,
// 		Index:  20,
// 		Events: []Event{},
// 	}
//
// 	if events.SetIndex(6) {
// 		t.Errorf("Incorrect return from SetIndex - expected:%v, got:%v", false, true)
// 	}
//
// 	if events.Index != 20 {
// 		t.Errorf("SetIndex incorrected updated internal index - expected:%v, got:%v", 20, events.Index)
// 	}
// }
