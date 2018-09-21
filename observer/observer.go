package observer

type Event struct {
	EventId   int
	Data      string
}

type IObserver interface {
	OnNotify(e Event)
	GetEventIds() []int
}

type ISubject interface {
	AddObserver()
	RemObserver()
	Notify(e Event)
}

