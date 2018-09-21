package factory

import (
	"design_patterns/adapter"
	"design_patterns/observer"
	"fmt"
)

type Monster struct {
	Name   string
	Height int32
	Gender int32
	Skin   interface{}
}

func (r Monster) Walk() {
	fmt.Println("monster walk")
}

func (r Monster) Run() {}

func (r Monster) Attack() {}

func (r Monster) ReleaseSkill(s adapter.Skill) {
	s.Release()
}

func (r Monster) SetGender() {}

func (r Monster) SetName() {}

func (r Monster) SetHeight() {}

func (r Monster) SetSkin() {}

func (r Monster) BeenAtacked() {
	//被击杀的话，发出事件通知
	obsMgr := observer.GetObsMgr()
	e := observer.Event{}
	e.EventId = observer.EventIdKillMonster
	e.Data = r.Name
	obsMgr.Notify(e)
}
