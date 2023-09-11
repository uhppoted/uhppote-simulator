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

func TestDoorInterlock1(t *testing.T) {
	tests := []struct {
		button      uint8
		open        map[uint8]bool
		interlocked bool
	}{
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},

		{1, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, true},
		{1, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false},
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false},

		{2, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, true},
		{2, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false},

		{3, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false},

		{4, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false},
	}

	for _, test := range tests {
		doors := MakeDoors()
		doors.Interlock = 1

		doors.doors[1].open = test.open[1]
		doors.doors[2].open = test.open[2]
		doors.doors[3].open = test.open[3]
		doors.doors[4].open = test.open[4]

		interlocked := doors.IsInterlocked(test.button)

		if interlocked != test.interlocked {
			t.Errorf("Unexpected 'IsInterlocked' -  expected:%v, got:%v", test.interlocked, interlocked)
		}
	}
}

func TestDoorPressButtonWithInterlock2(t *testing.T) {
	tests := []struct {
		button      uint8
		open        map[uint8]bool
		interlocked bool
	}{
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},

		{1, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false},
		{1, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false},
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false},

		{2, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false},
		{2, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false},

		{3, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, true},

		{4, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, true},
	}

	for _, test := range tests {
		doors := MakeDoors()
		doors.Interlock = 2

		doors.doors[1].open = test.open[1]
		doors.doors[2].open = test.open[2]
		doors.doors[3].open = test.open[3]
		doors.doors[4].open = test.open[4]

		interlocked := doors.IsInterlocked(test.button)

		if interlocked != test.interlocked {
			t.Errorf("Unexpected 'IsInterlocked' -  expected:%v, got:%v", test.interlocked, interlocked)
		}
	}
}

func TestDoorPressButtonWithInterlock3(t *testing.T) {
	tests := []struct {
		button      uint8
		open        map[uint8]bool
		interlocked bool
	}{
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},

		{1, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false},
		{1, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, true},
		{1, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false},
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false},

		{2, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, true},
		{2, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false},
		{2, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false},

		{3, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, true},

		{4, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, true},
		{4, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false},
	}

	for _, test := range tests {
		doors := MakeDoors()
		doors.Interlock = 3

		doors.doors[1].open = test.open[1]
		doors.doors[2].open = test.open[2]
		doors.doors[3].open = test.open[3]
		doors.doors[4].open = test.open[4]

		interlocked := doors.IsInterlocked(test.button)

		if interlocked != test.interlocked {
			t.Errorf("Unexpected 'IsInterlocked' -  expected:%v, got:%v", test.interlocked, interlocked)
		}
	}
}

func TestDoorPressButtonWithInterlock4(t *testing.T) {
	tests := []struct {
		button      uint8
		open        map[uint8]bool
		interlocked bool
	}{
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},

		{1, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false},
		{1, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, true},
		{1, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, true},
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false},

		{2, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, true},
		{2, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false},
		{2, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, true},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false},

		{3, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, true},
		{3, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, true},
		{3, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false},

		{4, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false},
	}

	for _, test := range tests {
		doors := MakeDoors()
		doors.Interlock = 4

		doors.doors[1].open = test.open[1]
		doors.doors[2].open = test.open[2]
		doors.doors[3].open = test.open[3]
		doors.doors[4].open = test.open[4]

		interlocked := doors.IsInterlocked(test.button)

		if interlocked != test.interlocked {
			t.Errorf("Unexpected 'IsInterlocked' -  expected:%v, got:%v", test.interlocked, interlocked)
		}
	}
}

func TestDoorPressButtonWithInterlock8(t *testing.T) {
	tests := []struct {
		button      uint8
		open        map[uint8]bool
		interlocked bool
	}{
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},
		{4, map[uint8]bool{1: false, 2: false, 3: false, 4: false}, false},

		{1, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, false},
		{1, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, true},
		{1, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, true},
		{1, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, true},

		{2, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, true},
		{2, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, false},
		{2, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, true},
		{2, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, true},

		{3, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, true},
		{3, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, true},
		{3, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, false},
		{3, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, true},

		{4, map[uint8]bool{1: true, 2: false, 3: false, 4: false}, true},
		{4, map[uint8]bool{1: false, 2: true, 3: false, 4: false}, true},
		{4, map[uint8]bool{1: false, 2: false, 3: true, 4: false}, true},
		{4, map[uint8]bool{1: false, 2: false, 3: false, 4: true}, false},
	}

	for _, test := range tests {
		doors := MakeDoors()
		doors.Interlock = 8

		doors.doors[1].open = test.open[1]
		doors.doors[2].open = test.open[2]
		doors.doors[3].open = test.open[3]
		doors.doors[4].open = test.open[4]

		interlocked := doors.IsInterlocked(test.button)

		if interlocked != test.interlocked {
			t.Errorf("Unexpected 'IsInterlocked' -  expected:%v, got:%v", test.interlocked, interlocked)
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
				Passcodes:    []uint32{12345, 0, 999999, 654321},
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

	expected := `{"interlock":3,"1":{"control":1,"delay":"5s","passcodes":[12345,0,999999,654321]},"2":{"control":2,"delay":"7.5s"},"3":{"control":3,"delay":"7s"},"4":{"control":3,"delay":"5s"}}`

	bytes, err := json.Marshal(doors)
	if err != nil {
		t.Fatalf("Error marshalling valid 'doors' struct (%v)", err)
	}

	if string(bytes) != expected {
		t.Errorf("Incorrectly marshalled 'doors' struct\n   expected:%v\n   got:     %v", expected, string(bytes))
	}
}

func TestDoorsUnmarshalJSON(t *testing.T) {
	bytes := []byte(`{"interlock":3,"1":{"control":1,"delay":"5s", "passcodes":[12345,0,999999,654321]},"2":{"control":2,"delay":"7.5s"},"3":{"control":3,"delay":"7s"},"4":{"control":3,"delay":"5s"}}`)

	doors := Doors{
		doors: map[uint8]*Door{},
	}

	expected := Doors{
		Interlock: 3,
		doors: map[uint8]*Door{
			1: &Door{
				ControlState: 1,
				Delay:        Delay(5 * time.Second),
				Passcodes:    []uint32{12345, 0, 999999, 654321},
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
