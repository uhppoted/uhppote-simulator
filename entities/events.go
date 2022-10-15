package entities

import (
	"encoding/json"
	"time"

	"github.com/uhppoted/uhppote-core/types"
)

type Event struct {
	Index     uint32          `json:"index"`
	Type      uint8           `json:"type"`
	Granted   bool            `json:"granted"`
	Door      uint8           `json:"door"`
	Direction uint8           `json:"direction"`
	Card      uint32          `json:"card"`
	Timestamp *types.DateTime `json:"timestamp,omitempty"`
	Reason    uint8           `json:"reason"`
}

// PENDING rework of uhpppote-core::DateTime.MarshalJSON
func (e Event) MarshalJSON() ([]byte, error) {
	event := struct {
		Index     uint32 `json:"index"`
		Type      uint8  `json:"type"`
		Granted   bool   `json:"granted"`
		Door      uint8  `json:"door"`
		Direction uint8  `json:"direction"`
		Card      uint32 `json:"card"`
		Timestamp string `json:"timestamp,omitempty"`
		Reason    uint8  `json:"reason"`
	}{
		Index:     e.Index,
		Type:      e.Type,
		Granted:   e.Granted,
		Door:      e.Door,
		Direction: e.Direction,
		Card:      e.Card,
		Reason:    e.Reason,
	}

	if e.Timestamp != nil {
		event.Timestamp = (*time.Time)(e.Timestamp).Format("2006-01-02 15:04:05")
	}

	return json.Marshal(event)
}

// PENDING rework of uhpppote-core::DateTime.UnmarshalJSON
func (e *Event) UnmarshalJSON(b []byte) error {
	event := struct {
		Index     uint32 `json:"index"`
		Type      uint8  `json:"type"`
		Granted   bool   `json:"granted"`
		Door      uint8  `json:"door"`
		Direction uint8  `json:"direction"`
		Card      uint32 `json:"card"`
		Timestamp string `json:"timestamp"`
		Reason    uint8  `json:"reason"`
	}{}

	if err := json.Unmarshal(b, &event); err != nil {
		return err
	}

	e.Index = event.Index
	e.Type = event.Type
	e.Granted = event.Granted
	e.Door = event.Door
	e.Direction = event.Direction
	e.Card = event.Card
	e.Reason = event.Reason

	if t, err := time.ParseInLocation("2006-01-02 15:04:05", event.Timestamp, time.Local); err == nil {
		timestamp := types.DateTime(time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.Local))
		e.Timestamp = &timestamp
	}

	return nil
}
