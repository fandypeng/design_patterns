package builder

import (
	"design_patterns/factory"
)

type IBuilder interface {
}

type builder struct {
	role factory.Behavior
}

func (b *builder) SetRole(r factory.Behavior) {
	b.role = r
}

func (b *builder) SetGender() *builder {
	b.role.SetGender()
	return b
}

func (b *builder) SetName() *builder {
	b.role.SetName()
	return b
}

func (b *builder) SetHeight() *builder {
	b.role.SetHeight()
	return b
}

func (b *builder) SetSkin() *builder {
	b.role.SetSkin()

	return b
}

func (b *builder) Build() {
	//build end
}
