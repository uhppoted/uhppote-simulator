package entities

import (
	"encoding/json"
	"math"
	"sort"
)

type EventList struct {
	size   uint32
	chunk  uint32
	index  uint32
	events []Event
}

const DEFAULT_EVENTLIST_SIZE = 256
const DEFAULT_EVENTLIST_CHUNK = 8

func NewEventList() EventList {
	return EventList{
		size:   DEFAULT_EVENTLIST_SIZE,
		chunk:  DEFAULT_EVENTLIST_CHUNK,
		index:  0,
		events: []Event{},
	}
}

func (l EventList) MarshalJSON() ([]byte, error) {
	list := struct {
		Size   uint32  `json:"size"`
		Chunk  uint32  `json:"chunk"`
		Index  uint32  `json:"index"`
		Events []Event `json:"events"`
	}{
		Size:   l.size,
		Chunk:  l.chunk,
		Index:  l.index,
		Events: l.events,
	}

	b, err := json.Marshal(list)

	return b, err
}

func (l *EventList) UnmarshalJSON(b []byte) error {
	list := struct {
		Size   uint32  `json:"size,omitempty"`
		Chunk  uint32  `json:"chunk,omitempty"`
		Index  uint32  `json:"index"`
		Events []Event `json:"events"`
	}{
		Size:  DEFAULT_EVENTLIST_SIZE,
		Chunk: DEFAULT_EVENTLIST_CHUNK,
	}

	if err := json.Unmarshal(b, &list); err != nil {
		return err
	}

	l.size = list.Size
	l.chunk = list.Chunk
	l.index = list.Index
	l.events = []Event{}

	if l.size == 0 {
		l.size = DEFAULT_EVENTLIST_SIZE
	}

	if l.chunk == 0 {
		l.chunk = DEFAULT_EVENTLIST_CHUNK
	}

	sort.SliceStable(list.Events, func(i, j int) bool { return list.Events[i].Index < list.Events[j].Index })

	index := uint32(0)
	for _, e := range list.Events {
		if e.Index > index {
			l.events = append(l.events, e)
			index = e.Index
		}
	}

	if N := len(l.events) - int(l.size); N > 0 {
		if l.chunk > 0 {
			n := float64(N)
			chunk := int(l.chunk)
			offset := chunk * int(math.Ceil(n/float64(chunk)))

			l.events = l.events[offset:]
		} else {
			l.events = l.events[N:]
		}
	}

	return nil
}

func (l *EventList) Add(event Event) uint32 {
	index := uint32(1)
	if N := len(l.events); N > 0 {
		index = l.events[N-1].Index + 1
	}

	event.Index = index

	l.events = append(l.events, event)
	if len(l.events) > int(l.size) {
		l.events = l.events[l.chunk:]
	}

	return index
}

func (l *EventList) Get(index uint32) Event {
	if N := len(l.events); N == 0 {
		return Event{}
	} else if index == 0 {
		return l.events[0]
	} else if index == 0xffffffff {
		return l.events[N-1]
	} else if index < l.events[0].Index {
		return Event{
			Type: 0xff,
		}
	} else if index > l.events[N-1].Index {
		return Event{}
	} else {
		for _, e := range l.events {
			if e.Index == index {
				return e
			}
		}
	}

	return Event{}
}

func (l *EventList) GetIndex() uint32 {
	return l.index
}

func (l *EventList) SetIndex(index uint32) bool {
	if index == l.index {
		return false
	}

	if index == 0 {
		l.index = index
		return true
	}

	if N := len(l.events); N > 0 {
		last := l.events[N-1].Index
		if index <= last {
			l.index = index
			return true
		}
	}

	return false
}

// NOTE: no validation - for unit tests only
func MakeEventList(index uint32, events []Event) EventList {
	return EventList{
		size:   256,
		chunk:  8,
		index:  index,
		events: events,
	}
}
