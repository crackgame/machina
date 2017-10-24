package machina

type Event interface {
	Client() Client
	FromState() string
	ToState() string
}

type BaseEvent struct {
	client    Client
	fromState string
	toState   string
}

func (e *BaseEvent) Client() Client {
	return e.client
}

func (e *BaseEvent) FromState() string {
	return e.fromState
}

func (e *BaseEvent) ToState() string {
	return e.toState
}

func NewBaseEvent(client Client, fromState string, toState string) Event {
	return &BaseEvent{
		client:    client,
		fromState: fromState,
		toState:   toState,
	}
}
