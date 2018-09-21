package decorator

import (
	"design_patterns/factory"
	"fmt"
)

type IDance interface {
	Dance()
}

type DanceMonster struct {
	factory.Behavior
}

func (dm DanceMonster) Dance() {
	fmt.Println("monster dance")
}
