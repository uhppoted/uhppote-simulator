package entities

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestDoorPressButton(t *testing.T) {
	doors := MakeDoors()

	ok, reason := doors.PressButton(1, 5*time.Second)

	if !ok {
		t.Errorf("Unexpected 'button press' fail -  expected:%v, got:%v", true, ok)
	}

	if reason != 0 {
		t.Errorf("Incorrect reason returned from button press - expected:%v, got:%v", 0, reason)
	}
}

func TestDoorPressButtonWithInterlock1(t *testing.T) {
	tests := []struct {
		button uint8
		open   map[uint8]bool
		ok     bool
		reason uint8
	}{
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},
		{4, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},

		{1, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false, ReasonInterlock},
		{1, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, true, ReasonOk},
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, true, ReasonOk},

		{2, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false, ReasonInterlock},
		{2, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, true, ReasonOk},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, true, ReasonOk},

		{3, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, true, ReasonOk},
		{3, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, true, ReasonOk},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false, ReasonInterlock},

		{4, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, true, ReasonOk},
		{4, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, true, ReasonOk},
		{4, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false, ReasonInterlock},
	}

	for _, test := range tests {
		doors := MakeDoors()
		doors.Interlock = 1

		doors.doors[1].open = test.open[1]
		doors.doors[2].open = test.open[2]
		doors.doors[3].open = test.open[3]
		doors.doors[4].open = test.open[4]

		ok, reason := doors.PressButton(test.button, 5*time.Second)

		if ok != test.ok {
			t.Errorf("Unexpected 'button press' ok -  expected:%v, got:%v", test.ok, ok)
		}

		if reason != test.reason {
			t.Errorf("Incorrect reason returned from button press - expected:%v, got:%v", test.reason, reason)
		}
	}
}

func TestDoorPressButtonWithInterlock2(t *testing.T) {
	tests := []struct {
		button uint8
		open   map[uint8]bool
		ok     bool
		reason uint8
	}{
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},
		{4, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},

		{1, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, true, ReasonOk},
		{1, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false, ReasonInterlock},
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, true, ReasonOk},

		{2, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, true, ReasonOk},
		{2, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, true, ReasonOk},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false, ReasonInterlock},

		{3, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false, ReasonInterlock},
		{3, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, true, ReasonOk},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, true, ReasonOk},

		{4, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, true, ReasonOk},
		{4, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false, ReasonInterlock},
		{4, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, true, ReasonOk},
	}

	for _, test := range tests {
		doors := MakeDoors()
		doors.Interlock = 2

		doors.doors[1].open = test.open[1]
		doors.doors[2].open = test.open[2]
		doors.doors[3].open = test.open[3]
		doors.doors[4].open = test.open[4]

		ok, reason := doors.PressButton(test.button, 5*time.Second)

		if ok != test.ok {
			t.Errorf("Unexpected 'button press' ok -  expected:%v, got:%v", test.ok, ok)
		}

		if reason != test.reason {
			t.Errorf("Incorrect reason returned from button press - expected:%v, got:%v", test.reason, reason)
		}
	}
}

func TestDoorPressButtonWithInterlock3(t *testing.T) {
	tests := []struct {
		button uint8
		open   map[uint8]bool
		ok     bool
		reason uint8
	}{
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},
		{4, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},

		{1, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false, ReasonInterlock},
		{1, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false, ReasonInterlock},
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, true, ReasonOk},

		{2, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false, ReasonInterlock},
		{2, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false, ReasonInterlock},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, true, ReasonOk},

		{3, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false, ReasonInterlock},
		{3, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false, ReasonInterlock},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, true, ReasonOk},

		{4, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, true, ReasonOk},
		{4, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, true, ReasonOk},
		{4, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, true, ReasonOk},
	}

	for _, test := range tests {
		doors := MakeDoors()
		doors.Interlock = 3

		doors.doors[1].open = test.open[1]
		doors.doors[2].open = test.open[2]
		doors.doors[3].open = test.open[3]
		doors.doors[4].open = test.open[4]

		ok, reason := doors.PressButton(test.button, 5*time.Second)

		if ok != test.ok {
			t.Errorf("Unexpected 'button press' ok -  expected:%v, got:%v", test.ok, ok)
		}

		if reason != test.reason {
			t.Errorf("Incorrect reason returned from button press - expected:%v, got:%v", test.reason, reason)
		}
	}
}

func TestDoorPressButtonWithInterlock4(t *testing.T) {
	tests := []struct {
		button uint8
		open   map[uint8]bool
		ok     bool
		reason uint8
	}{
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},
		{4, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, true, ReasonOk},

		{1, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false, ReasonInterlock},
		{1, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false, ReasonInterlock},
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false, ReasonInterlock},

		{2, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false, ReasonInterlock},
		{2, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false, ReasonInterlock},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false, ReasonInterlock},

		{3, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false, ReasonInterlock},
		{3, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false, ReasonInterlock},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false, ReasonInterlock},

		{4, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false, ReasonInterlock},
		{4, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false, ReasonInterlock},
		{4, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false, ReasonInterlock},
	}

	for _, test := range tests {
		doors := MakeDoors()
		doors.Interlock = 4

		doors.doors[1].open = test.open[1]
		doors.doors[2].open = test.open[2]
		doors.doors[3].open = test.open[3]
		doors.doors[4].open = test.open[4]

		ok, reason := doors.PressButton(test.button, 5*time.Second)

		if ok != test.ok {
			t.Errorf("Unexpected 'button press' ok -  expected:%v, got:%v", test.ok, ok)
		}

		if reason != test.reason {
			t.Errorf("Incorrect reason returned from button press - expected:%v, got:%v", test.reason, reason)
		}
	}
}

func TestDoorsMarshalJSON(t *testing.T) {
	doors := Doors{
		Interlock: 3,
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

	expected := `{"interlock":3,"1":{"control":1,"delay":"5s"},"2":{"control":2,"delay":"7.5s"},"3":{"control":3,"delay":"7s"},"4":{"control":3,"delay":"5s"}}`

	bytes, err := json.Marshal(doors)
	if err != nil {
		t.Fatalf("Error marshalling valid 'doors' struct (%v)", err)
	}

	if string(bytes) != expected {
		t.Errorf("Incorrectly marshalled 'doors' struct\n   expected:%v\n   got:     %v", expected, string(bytes))
	}
}

func TestDoorsUnmarshalJSON(t *testing.T) {
	bytes := []byte(`{"interlock":3,"1":{"control":1,"delay":"5s"},"2":{"control":2,"delay":"7.5s"},"3":{"control":3,"delay":"7s"},"4":{"control":3,"delay":"5s"}}`)

	doors := Doors{
		doors: map[uint8]*Door{},
	}

	expected := Doors{
		Interlock: 3,
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

func TestDoorsUnmarshalInvalidJSON(t *testing.T) {
	bytes := []byte(`{"interlock":3,"1":{"control":0,"delay":"5s"},"2":{"control":4,"delay":"7.5s"},"3":{"control":5,"delay":"7s"},"4":{"control":6,"delay":"5s"}}`)

	doors := Doors{
		doors: map[uint8]*Door{},
	}

	expected := Doors{
		Interlock: 3,
		doors: map[uint8]*Door{
			1: &Door{
				ControlState: 3,
				Delay:        Delay(5 * time.Second),
			},
			2: &Door{
				ControlState: 3,
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
		t.Fatalf("Error unmarshalling 'doors' JSON with invalid control state (%v)", err)
	}

	if !reflect.DeepEqual(doors, expected) {
		t.Errorf("Incorrectly unmarshalled 'doors' struct\n   expected:%#v\n   got:     %#v", expected, doors)
	}
}
