package factory

import (
	"design_patterns/adapter"
	"fmt"
)

type Player struct {
	Name   string
	Height int32
	Gender int32
	Skin   interface{}
}

func (r Player) Walk() {
	fmt.Println("player walk")
}

func (r Player) Run() {}

func (r Player) Attack() {}

func (r Player) ReleaseSkill(s adapter.Skill) {
	s.Release()
}

func (r Player) BeenAtacked() {

}

func (r Player) SetGender() {}

func (r Player) SetName() {}

func (r Player) SetHeight() {}

func (r Player) SetSkin() {}
