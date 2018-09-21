package factory

import "design_patterns/adapter"

type Property interface {
	SetGender()
	SetName()
	SetHeight()
	SetSkin()
}

type Behavior interface {
	Property
	Walk()
	Run()
	Attack()
	ReleaseSkill(s adapter.Skill)
	BeenAtacked()
}
