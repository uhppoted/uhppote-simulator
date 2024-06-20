package entities

import (
	"encoding/json"
	"testing"
	"time"

	lib "github.com/uhppoted/uhppote-core/types"
)

func TestTaskListToJSON(t *testing.T) {
	tasklist := TaskList{
		Tasks: []lib.Task{
			lib.Task{
				Task: lib.DoorControlled,
				Door: 2,
				From: lib.MustParseDate("2024-01-01"),
				To:   lib.MustParseDate("2024-12-31"),
				Weekdays: lib.Weekdays{
					time.Monday:    true,
					time.Wednesday: false,
					time.Friday:    true,
				},
				Start: lib.NewHHmm(12, 34),
				Cards: 17,
			},
		},
		added: []lib.Task{
			lib.Task{
				Task: lib.EnableTimeProfile,
				Door: 3,
				From: lib.MustParseDate("2023-01-01"),
				To:   lib.MustParseDate("2025-12-31"),
				Weekdays: lib.Weekdays{
					time.Tuesday: true,
				},
				Start: lib.NewHHmm(21, 22),
				Cards: 13,
			},
		},
		scheduled: map[int]scheduled{},
		last:      "2024-06-20",
	}

	expected := `{"tasks":[{"task":"CONTROL DOOR","door":2,"start-date":"2024-01-01","end-date":"2024-12-31","weekdays":"Monday,Friday","start":"12:34","cards":17}]}`

	if bytes, err := json.Marshal(tasklist); err != nil {
		t.Fatalf("Error marshalling task list (%v)", err)
	} else if string(bytes) != expected {
		t.Errorf("TaskList marshalled incorrectly:\n   expected:%v\n   got:     %v", expected, string(bytes))
	}
}

func TestEmptyTaskListToJSON(t *testing.T) {
	tasklist := TaskList{
		Tasks: []lib.Task{},
		added: []lib.Task{
			lib.Task{
				Task: lib.EnableTimeProfile,
				Door: 3,
				From: lib.MustParseDate("2023-01-01"),
				To:   lib.MustParseDate("2025-12-31"),
				Weekdays: lib.Weekdays{
					time.Tuesday: true,
				},
				Start: lib.NewHHmm(21, 22),
				Cards: 13,
			},
		},
		scheduled: map[int]scheduled{},
		last:      "2024-06-20",
	}

	expected := `{}`

	if bytes, err := json.Marshal(tasklist); err != nil {
		t.Fatalf("Error marshalling task list (%v)", err)
	} else if string(bytes) != expected {
		t.Errorf("TaskList marshalled incorrectly:\n   expected:%v\n   got:     %v", expected, string(bytes))
	}
}
