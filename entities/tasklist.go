package entities

import (
	"github.com/uhppoted/uhppote-core/types"
)

type TaskList struct {
	Tasks []types.Task `json:"tasks"`
	added []types.Task
}

func (t *TaskList) Add(task types.Task) bool {
	t.added = append(t.added, task)

	return true
}

func (t *TaskList) Clear() bool {
	t.Tasks = []types.Task{}
	t.added = []types.Task{}

	return true
}

func (t *TaskList) Refresh() bool {
	t.Tasks = append(t.Tasks, t.added...)
	t.added = []types.Task{}

	return true
}
