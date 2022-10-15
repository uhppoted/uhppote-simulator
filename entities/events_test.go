package entities

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/uhppoted/uhppote-core/types"
)

func TestMarshalEvent(t *testing.T) {
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 9, time.Local))

	e := Event{
		Index:     7,
		Type:      13,
		Granted:   true,
		Door:      3,
		Direction: 1,
		Card:      8165535,
		Timestamp: &timestamp,
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
	timestamp := types.DateTime(time.Date(2021, time.December, 27, 13, 14, 15, 0, time.Local))

	expected := Event{
		Index:     7,
		Type:      13,
		Granted:   true,
		Door:      3,
		Direction: 1,
		Card:      8165535,
		Timestamp: &timestamp,
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
