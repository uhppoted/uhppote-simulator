package entities

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/uhppoted/uhppote-core/types"
)

func TestSetAntiPassback(t *testing.T) {
	antipassback := AntiPassback{
		antipassback: types.Readers12_34,
		deny: []pair{
			{10058400, 1},
			{10058400, 3},
			{10058399, 2},
			{10058399, 4},
		},
	}

	expected := AntiPassback{
		antipassback: types.Readers1_234,
		deny:         []pair{},
	}

	if ok := antipassback.Set(0x04); !ok {
		t.Errorf("incorrect return value:%v, got:%v", true, ok)
	} else if !reflect.DeepEqual(antipassback, expected) {
		t.Errorf("incorrect entity\n   expected: %v\ngot:    %v", expected, antipassback)
	}
}

func TestAntiPassbackDisabled(t *testing.T) {
	antipassback := AntiPassback{
		antipassback: types.Disabled,
	}

	sequence := []struct {
		card     uint32
		door     uint8
		expected bool
	}{
		{10058400, 1, true},
		{10058400, 1, true},
		{10058400, 2, true},
		{10058400, 2, true},
		{10058400, 3, true},
		{10058400, 3, true},
		{10058400, 4, true},
		{10058400, 4, true},
	}

	for _, v := range sequence {
		if allowed := antipassback.Allow(v.card, v.door); allowed != v.expected {
			t.Fatalf("incorrect 'deny' - expected:%v, got:%v", v.expected, allowed)
		} else if allowed {
			antipassback.Allowed(v.card, v.door)
		}
	}
}

func TestAntiPassback12_34(t *testing.T) {
	antipassback := AntiPassback{
		antipassback: types.Readers12_34,
	}

	sequence := []struct {
		card     uint32
		door     uint8
		expected bool
	}{
		{10058400, 1, true},
		{10058400, 1, false},
		{10058400, 1, false},
		{10058400, 2, true},
		{10058400, 2, false},
		{10058400, 2, false},
		{10058400, 1, true},
		{10058400, 1, false},
		{10058400, 1, false},

		{10058400, 3, true},
		{10058400, 3, false},
		{10058400, 3, false},
		{10058400, 4, true},
		{10058400, 4, false},
		{10058400, 4, false},
		{10058400, 3, true},
		{10058400, 3, false},
		{10058400, 3, false},
	}

	for _, v := range sequence {
		if allowed := antipassback.Allow(v.card, v.door); allowed != v.expected {
			t.Fatalf("incorrect 'deny' - expected:%v, got:%v", v.expected, allowed)
		} else if allowed {
			antipassback.Allowed(v.card, v.door)
		}
	}
}

func TestAntiPassback13_24(t *testing.T) {
	antipassback := AntiPassback{
		antipassback: types.Readers13_24,
	}

	sequence := []struct {
		card     uint32
		door     uint8
		expected bool
	}{
		{10058400, 1, true},
		{10058400, 1, false},
		{10058400, 1, false},
		{10058400, 3, true},
		{10058400, 3, false},
		{10058400, 3, false},
		{10058400, 1, true},
		{10058400, 1, false},
		{10058400, 1, false},

		{10058400, 2, true},
		{10058400, 2, false},
		{10058400, 2, false},
		{10058400, 4, true},
		{10058400, 4, false},
		{10058400, 4, false},
		{10058400, 2, true},
		{10058400, 2, false},
		{10058400, 2, false},
	}

	for _, v := range sequence {
		if allowed := antipassback.Allow(v.card, v.door); allowed != v.expected {
			t.Fatalf("incorrect 'deny' - expected:%v, got:%v", v.expected, allowed)
		} else if allowed {
			antipassback.Allowed(v.card, v.door)
		}
	}
}

func TestAntiPassback1_23(t *testing.T) {
	antipassback := AntiPassback{
		antipassback: types.Readers1_23,
	}

	sequence := []struct {
		card     uint32
		door     uint8
		expected bool
	}{
		{10058400, 1, true},
		{10058400, 1, false},
		{10058400, 2, true},
		{10058400, 2, false},
		{10058400, 3, false},
		{10058400, 1, true},
		{10058400, 2, true},

		{10058400, 1, true},
		{10058400, 1, false},
		{10058400, 3, true},
		{10058400, 2, false},
		{10058400, 3, false},
		{10058400, 1, true},
		{10058400, 3, true},

		{10058400, 4, true},
		{10058400, 4, true},
	}

	for _, v := range sequence {
		if allowed := antipassback.Allow(v.card, v.door); allowed != v.expected {
			t.Fatalf("incorrect 'deny' - expected:%v, got:%v", v.expected, allowed)
		} else if allowed {
			antipassback.Allowed(v.card, v.door)
		}
	}
}

func TestAntiPassback1_234(t *testing.T) {
	antipassback := AntiPassback{
		antipassback: types.Readers1_234,
	}

	sequence := []struct {
		card     uint32
		door     uint8
		expected bool
	}{
		{10058400, 1, true},
		{10058400, 1, false},
		{10058400, 2, true},
		{10058400, 2, false},
		{10058400, 3, false},
		{10058400, 4, false},
		{10058400, 1, true},
		{10058400, 2, true},

		{10058400, 1, true},
		{10058400, 1, false},
		{10058400, 3, true},
		{10058400, 2, false},
		{10058400, 3, false},
		{10058400, 4, false},
		{10058400, 1, true},
		{10058400, 3, true},

		{10058400, 1, true},
		{10058400, 1, false},
		{10058400, 4, true},
		{10058400, 2, false},
		{10058400, 3, false},
		{10058400, 4, false},
		{10058400, 1, true},
		{10058400, 4, true},
	}

	for _, v := range sequence {
		if allowed := antipassback.Allow(v.card, v.door); allowed != v.expected {
			t.Fatalf("incorrect 'deny' - expected:%v, got:%v", v.expected, allowed)
		} else if allowed {
			antipassback.Allowed(v.card, v.door)
		}
	}
}

func TestAntiPassbackMarshal(t *testing.T) {
	antipassback := AntiPassback{
		antipassback: types.Readers1_234,
	}

	expected := "4"

	if encoded, err := json.Marshal(antipassback); err != nil {
		t.Fatalf("error marshalling AntiPassback (%v)", err)
	} else if string(encoded) != expected {
		t.Errorf("incorrectly marshalled JSON\n   expected:%v\n   got:     %v", expected, string(encoded))
	}
}

func TestAntiPassbackUnmarshal(t *testing.T) {
	expected := AntiPassback{
		antipassback: types.Readers1_234,
	}

	encoded := []byte("4")
	antipassback := AntiPassback{}

	if err := json.Unmarshal(encoded, &antipassback); err != nil {
		t.Fatalf("error unmarshalling AntiPassback (%v)", err)
	} else if !reflect.DeepEqual(antipassback, expected) {
		t.Errorf("incorrectly unmarshalled AntiPassback\n   expected:%v\n   got:     %v", expected, antipassback)
	}
}
