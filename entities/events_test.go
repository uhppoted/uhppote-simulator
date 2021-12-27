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
		First: 3,
		Last:  27,
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
  "first": 3,
  "last": 27,
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
  "first": 3,
  "last": 27,
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
		First: 3,
		Last:  27,
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

func TestUnmarshalEventListWithDefaultSizeAndChunk(t *testing.T) {
	timestamp := time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local)

	bytes := []byte(`{
  "first": 3,
  "last": 27,
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
		First: 3,
		Last:  27,
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

func TestSetIndexWithZero(t *testing.T) {
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
