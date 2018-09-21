package observer

import (
	"container/list"
	"sync"
)

type MObserver struct {
	obs map[int]*list.List
}

var singleObsMgr *MObserver
var once sync.Once

func GetObsMgr() *MObserver {
	once.Do(func() {
		singleObsMgr = new(MObserver)
		singleObsMgr.obs = make(map[int]*list.List)
	})
	return singleObsMgr
}

func (m *MObserver) AddObserver(o IObserver) {
	for _, eid := range o.GetEventIds() {
		if _, ok := m.obs[eid]; !ok {
			m.obs[eid] = list.New()
		}
		m.obs[eid].PushBack(o)
	}
}

func (m *MObserver) RemObserver(o IObserver) {
	for _, eid := range o.GetEventIds() {
		if _, ok := m.obs[eid]; !ok {
			continue
		}
		for ob := m.obs[eid].Front(); ob != nil; ob = ob.Next() {
			if ob.Value.(*IObserver) == &o {
				m.obs[eid].Remove(ob)
				break
			}
		}
	}
}

func (m *MObserver) Notify(e Event) {
	eid := e.EventId
	if _, ok := m.obs[eid]; !ok {
		return
	}
	for ob := m.obs[eid].Front(); ob != nil; ob = ob.Next() {
		ob.Value.(IObserver).OnNotify(e)
	}
}
