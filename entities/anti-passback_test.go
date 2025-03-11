package entities

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/uhppoted/uhppote-core/types"
)

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
