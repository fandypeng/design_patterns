package observer

import "fmt"

type Achievement struct {
	eventIds []int
}

func NewAchievement() *Achievement {
	t := new(Achievement)
	t.eventIds = append(t.eventIds, EventIdKillMonster)
	t.eventIds = append(t.eventIds, EventIdLearnSkill)
	return t
}

func (t *Achievement) GetEventIds() []int {
	return t.eventIds
}

func (t *Achievement) OnNotify(e Event) {
	switch e.EventId {
	case EventIdKillMonster:
		fmt.Println("make an achievement kill a monster" + e.Data)
	case EventIdLearnSkill:
		fmt.Println("make an achievement learn new skill" + e.Data)
	}
}
