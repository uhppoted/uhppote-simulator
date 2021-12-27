package entities

import (
	"encoding/json"
	"time"

	"github.com/uhppoted/uhppote-core/types"
)

type Event struct {
	Index     uint32         `json:"index"`
	Type      uint8          `json:"type"`
	Granted   bool           `json:"granted"`
	Door      uint8          `json:"door"`
	Direction uint8          `json:"direction"`
	Card      uint32         `json:"card"`
	Timestamp types.DateTime `json:"timestamp"`
	Reason    uint8          `json:"reason"`
}

type EventList struct {
	Size   uint32  `json:"size"`
	Chunk  uint32  `json:"chunk"`
	First  uint32  `json:"first"`
	Last   uint32  `json:"last"`
	Index  uint32  `json:"index"`
	Events []Event `json:"events"`
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
		Timestamp string `json:"timestamp"`
		Reason    uint8  `json:"reason"`
	}{
		Index:     e.Index,
		Type:      e.Type,
		Granted:   e.Granted,
		Door:      e.Door,
		Direction: e.Direction,
		Card:      e.Card,
		Timestamp: time.Time(e.Timestamp).Format("2006-01-02 15:04:05"),
		Reason:    e.Reason,
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

	if t, err := time.ParseInLocation("2006-01-02 15:04:05", event.Timestamp, time.Local); err != nil {
		return err
	} else {
		e.Timestamp = types.DateTime(time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.Local))
	}

	return nil
}

func (l EventList) MarshalJSON() ([]byte, error) {
	list := struct {
		Size   uint32  `json:"size"`
		Chunk  uint32  `json:"chunk"`
		First  uint32  `json:"first"`
		Last   uint32  `json:"last"`
		Index  uint32  `json:"index"`
		Events []Event `json:"events"`
	}{
		Size:   l.Size,
		Chunk:  l.Chunk,
		First:  l.First,
		Last:   l.Last,
		Index:  l.Index,
		Events: l.Events,
	}

	b, err := json.Marshal(list)

	return b, err
}

func (l *EventList) UnmarshalJSON(b []byte) error {
	list := struct {
		Size   uint32  `json:"size"`
		Chunk  uint32  `json:"chunk"`
		First  uint32  `json:"first"`
		Last   uint32  `json:"last"`
		Index  uint32  `json:"index"`
		Events []Event `json:"events"`
	}{
		Size:  64,
		Chunk: 8,
	}

	if err := json.Unmarshal(b, &list); err != nil {
		return err
	}

	l.Size = list.Size
	l.Chunk = list.Chunk
	l.First = list.First
	l.Last = list.Last
	l.Index = list.Index
	l.Events = list.Events

	return nil
}

func (l *EventList) Add(event *Event) {
	if event != nil {
		l.Last = l.Last + 1
		if l.Last > l.Size {
			l.Last = 1
		}

		if l.Last == l.First {
			l.First = l.First + 1
			if l.First > l.Size {
				l.First = 1
			}
		}

		event.Index = uint32(l.Last)

		index := l.Last
		if index >= uint32(len(l.Events)) {
			l.Events = append(l.Events, *event)
		} else {
			l.Events[index-1] = *event
		}
	}
}

func (l *EventList) Get(index uint32) *Event {
	if len(l.Events) > 0 {
		if index == 0 {
			return &l.Events[l.First-1]
		}

		if index == 0xffffffff || index > uint32(len(l.Events)) {
			return &l.Events[l.Last-1]
		}

		if index > 0 && int(index) <= len(l.Events) {
			return &l.Events[index-1]
		}
	}

	return nil
}

func (l *EventList) SetIndex(index uint32) bool {
	if index == l.Index {
		return false
	}

	if index == 0 {
		l.Index = index
		return true
	}

	if l.Last >= l.First {
		if index > l.Last || index < l.First {
			return false
		} else {
			l.Index = index
			return true
		}
	}

	// Events list has rolled over
	// FIXME verify with actual controller
	if index <= l.Last || index >= l.First {
		if index <= l.Size {
			l.Index = index
		} else {
			l.Index = 1
		}

		return true
	}

	return false
}
