package machina

type Client interface {
	State() string
	SetState(state string)

	GetLastEvent() string
	SetLastEvent(evtName string)

	GetLastPars() []interface{}
	SetLastPars(pars []interface{})
}

type BaseClient struct {
	state string

	lastEvent string
	lastPars  []interface{}
}

func (c *BaseClient) State() string {
	return c.state
}

func (c *BaseClient) SetState(state string) {
	c.state = state
}

func (c *BaseClient) GetLastEvent() string {
	return c.lastEvent
}

func (c *BaseClient) SetLastEvent(evtName string) {
	c.lastEvent = evtName
}

func (c *BaseClient) GetLastPars() []interface{} {
	return c.lastPars
}

func (c *BaseClient) SetLastPars(pars []interface{}) {
	c.lastPars = pars
}
