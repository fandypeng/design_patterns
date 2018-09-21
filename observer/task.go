package observer

import "fmt"

type Task struct {
	eventIds []int
}

func NewTask() *Task {
	t := new(Task)
	t.eventIds = append(t.eventIds, EventIdKillMonster)
	t.eventIds = append(t.eventIds, EventIdLearnSkill)
	return t
}

func (t *Task) GetEventIds() []int {
	return t.eventIds
}

func (t *Task) OnNotify(e Event) {
	switch e.EventId {
	case EventIdKillMonster:
		fmt.Println("finish task kill a monster" + e.Data)
	case EventIdLearnSkill:
		fmt.Println("finish task learn new skill" + e.Data)
	}
}
