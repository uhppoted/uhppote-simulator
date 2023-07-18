package entities

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/uhppoted/uhppote-core/types"
)

type TaskList struct {
	Tasks     []types.Task `json:"tasks"`
	added     []types.Task
	scheduled map[int]scheduled
	last      string
	guard     sync.Mutex
}

type scheduled struct {
	index int
	task  types.Task
}

func (t *TaskList) Add(task types.Task) bool {
	t.guard.Lock()
	defer t.guard.Unlock()

	t.added = append(t.added, task)

	return true
}

func (t *TaskList) Clear() bool {
	t.guard.Lock()
	defer t.guard.Unlock()

	t.Tasks = []types.Task{}
	t.added = []types.Task{}
	t.scheduled = map[int]scheduled{}

	return true
}

func (t *TaskList) Refresh() bool {
	t.guard.Lock()
	defer t.guard.Unlock()

	t.Tasks = append(t.Tasks, t.added...)
	t.added = []types.Task{}

	t.refresh(time.Now())

	return true
}

/** Expects the RWMutex to be locked/unlocked by the invoking function.
*
 */
func (t *TaskList) refresh(from time.Time) {
	now := types.HHmmFromTime(from)
	today := types.Date(time.Now())
	list := map[int]scheduled{}

	for ix, task := range t.Tasks {
		if task.From.After(today) || task.To.Before(today) || !task.Weekdays[today.Weekday()] {
			continue
		}

		if !task.Start.Before(now) {
			list[ix] = scheduled{
				index: ix,
				task:  task,
			}
		}
	}

	t.scheduled = list
}

func startOfDay() time.Time {
	today := time.Now()

	return time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
}

func (t *TaskList) Run(handler func(door uint8, task types.TaskType)) {
	t.guard.Lock()
	defer t.guard.Unlock()

	pending := []scheduled{}
	now := types.HHmmFromTime(time.Now())
	today := types.Date(time.Now())

	if t.last != fmt.Sprintf("%v", today) {
		t.refresh(startOfDay())
		t.last = fmt.Sprintf("%v", today)
	}

	for _, task := range t.scheduled {
		if !task.task.Start.After(now) {
			pending = append(pending, task)
		}
	}

	sort.SliceStable(pending, func(i, j int) bool {
		if pending[i].task.Start.Before(pending[j].task.Start) {
			return true
		} else {
			return pending[i].index < pending[j].index
		}
	})

	for _, p := range pending {
		handler(p.task.Door, p.task.Task)
	}

	for _, p := range pending {
		delete(t.scheduled, p.index)
	}
}
