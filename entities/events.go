package entities

import (
	"encoding/json"
	"sort"
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

type EventList struct {
	Size   uint32  `json:"size"`
	Chunk  uint32  `json:"chunk"`
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

func (l EventList) MarshalJSON() ([]byte, error) {
	list := struct {
		Size   uint32  `json:"size"`
		Chunk  uint32  `json:"chunk"`
		Index  uint32  `json:"index"`
		Events []Event `json:"events"`
	}{
		Size:   l.Size,
		Chunk:  l.Chunk,
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
		Index  uint32  `json:"index"`
		Events []Event `json:"events"`
	}{
		Size:  256,
		Chunk: 8,
	}

	if err := json.Unmarshal(b, &list); err != nil {
		return err
	}

	l.Size = list.Size
	l.Chunk = list.Chunk
	l.Index = list.Index
	l.Events = []Event{}

	sort.SliceStable(list.Events, func(i, j int) bool { return list.Events[i].Index < list.Events[j].Index })

	index := uint32(0)
	for _, e := range list.Events {
		if e.Index > index {
			l.Events = append(l.Events, e)
			index = e.Index
		}
	}

	for len(l.Events) > int(l.Size) {
		l.Events = l.Events[l.Chunk:]
	}

	return nil
}

func (l *EventList) Add(event Event) uint32 {
	index := uint32(1)
	if N := len(l.Events); N > 0 {
		index = l.Events[N-1].Index + 1
	}

	event.Index = index

	l.Events = append(l.Events, event)
	if len(l.Events) > int(l.Size) {
		l.Events = l.Events[l.Chunk:]
	}

	return index
}

func (l *EventList) Get(index uint32) Event {
	if N := len(l.Events); N == 0 {
		return Event{}
	} else if index == 0 {
		return l.Events[0]
	} else if index == 0xffffffff {
		return l.Events[N-1]
	} else if index < l.Events[0].Index {
		return Event{
			Type: 0xff,
		}
	} else if index > l.Events[N-1].Index {
		return Event{}
	} else {
		for _, e := range l.Events {
			if e.Index == index {
				return e
			}
		}
	}

	return Event{}
}

func (l *EventList) SetIndex(index uint32) bool {
	if index == l.Index {
		return false
	}

	if index == 0 {
		l.Index = index
		return true
	}

	if N := len(l.Events); N > 0 {
		last := l.Events[N-1].Index
		if index <= last {
			l.Index = index
			return true
		}
	}

	return false
}
