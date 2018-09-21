package adapter

type Skill interface {
	Before()
	Release()
	After()
}

