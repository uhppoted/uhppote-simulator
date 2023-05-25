package entities

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestDoorsMarshalJSON(t *testing.T) {
	doors := Doors{
		doors: map[uint8]*Door{
			1: &Door{
				ControlState: 1,
				Delay:        Delay(5 * time.Second),
			},
			2: &Door{
				ControlState: 2,
				Delay:        Delay(7500 * time.Millisecond),
			},
			3: &Door{
				ControlState: 3,
				Delay:        Delay(7 * time.Second),
			},
			4: &Door{
				ControlState: 3,
				Delay:        Delay(5 * time.Second),
			},
		},
	}

	expected := `{"doors":{"1":{"control":1,"delay":"5s"},"2":{"control":2,"delay":"7.5s"},"3":{"control":3,"delay":"7s"},"4":{"control":3,"delay":"5s"}}}`

	bytes, err := json.Marshal(doors)
	if err != nil {
		t.Fatalf("Error marshalling valid 'doors' struct (%v)", err)
	}

	if string(bytes) != expected {
		t.Errorf("Incorrectly marshalled 'doors' struct\n   expected:%v\n   got:     %v", expected, string(bytes))
	}
}

func TestDoorsUnmarshalJSON(t *testing.T) {
	bytes := []byte(`{"doors":{"1":{"control":1,"delay":"5s"},"2":{"control":2,"delay":"7.5s"},"3":{"control":3,"delay":"7s"},"4":{"control":3,"delay":"5s"}}}`)

	doors := Doors{
		doors: map[uint8]*Door{},
	}

	expected := Doors{
		doors: map[uint8]*Door{
			1: &Door{
				ControlState: 1,
				Delay:        Delay(5 * time.Second),
			},
			2: &Door{
				ControlState: 2,
				Delay:        Delay(7500 * time.Millisecond),
			},
			3: &Door{
				ControlState: 3,
				Delay:        Delay(7 * time.Second),
			},
			4: &Door{
				ControlState: 3,
				Delay:        Delay(5 * time.Second),
			},
		},
	}

	if err := json.Unmarshal(bytes, &doors); err != nil {
		t.Fatalf("Error unmarshalling valid 'doors' JSON (%v)", err)
	}

	if !reflect.DeepEqual(doors, expected) {
		t.Errorf("Incorrectly unmarshalled 'doors' struct\n   expected:%#v\n   got:     %#v", expected, doors)
	}
}

// TODO check unmarshal invalid control state
