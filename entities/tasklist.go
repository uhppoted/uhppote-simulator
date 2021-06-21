package entities

import (
	"sort"
	"time"

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

func (t *TaskList) Run(handler func(door uint8, task types.TaskType)) {
	tasks := []types.Task{}

	now := types.HHmmFromTime(time.Now())
	today := types.Date(time.Now())

	for _, task := range t.Tasks {
		if !task.From.After(today) && !task.To.Before(today) && !task.Start.After(now) && task.Weekdays[today.Weekday()] {
			tasks = append(tasks, task)
		}
	}

	sort.SliceStable(tasks, func(i, j int) bool { return tasks[i].Start.Before(tasks[j].Start) })

	doors := map[uint8]types.TaskType{}
	profiles := map[uint8]types.TaskType{}
	buttons := map[uint8]types.TaskType{}
	other := map[uint8]types.TaskType{}

	for _, task := range tasks {
		switch task.Task {
		case types.DoorControlled:
			doors[task.Door] = types.DoorControlled

		case types.DoorNormallyOpen:
			doors[task.Door] = types.DoorNormallyOpen

		case types.DoorNormallyClosed:
			doors[task.Door] = types.DoorNormallyClosed

		case types.DisableTimeProfile:
			profiles[task.Door] = types.DisableTimeProfile

		case types.EnableTimeProfile:
			profiles[task.Door] = types.EnableTimeProfile

			//	case types.CardNoPassword:
			//	case types.CardInPassword:
			//	case types.CardInOutPassword:
			//	case types.EnableMoreCards:
			//	case types.DisableMoreCards:

		case types.TriggerOnce:
			if task.Start.Equals(now) {
				other[task.Door] = types.TriggerOnce
			}

		case types.DisablePushButton:
			buttons[task.Door] = types.DisablePushButton

		case types.EnablePushButton:
			buttons[task.Door] = types.EnablePushButton
		}
	}

	for _, d := range []uint8{1, 2, 3, 4} {
		if v, ok := doors[d]; ok {
			handler(d, v)
		}
	}

	for _, d := range []uint8{1, 2, 3, 4} {
		if v, ok := profiles[d]; ok {
			handler(d, v)
		}
	}

	for _, d := range []uint8{1, 2, 3, 4} {
		if v, ok := buttons[d]; ok {
			handler(d, v)
		}
	}

	for _, d := range []uint8{1, 2, 3, 4} {
		if v, ok := other[d]; ok {
			handler(d, v)
		}
	}
}
