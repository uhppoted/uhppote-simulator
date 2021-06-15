package entities

import (
	"github.com/uhppoted/uhppote-core/types"
)

type TaskList []types.Task

func (t *TaskList) Add(task types.Task) bool {
	return false
}

func (t *TaskList) Clear() bool {
	*t = []types.Task{}

	return true
}
