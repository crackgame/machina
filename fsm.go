package machina

import "github.com/crackgame/emitter"

const (
	EvtTransition = "transition"
	EvtEnter      = "_onEnter"
	EvtExit       = "_onExit"
	Undefined     = "undefined"
	AlwayEvent    = "*"
)

type FSM struct {
	initialState string
	states       map[string]*State
	*emitter.Emitter
}

func NewFSM(initialState string, states map[string]*State) *FSM {
	fsm := &FSM{
		initialState: initialState,
		states:       states,
		Emitter:      emitter.NewEmitter(),
	}
	return fsm
}

func (f *FSM) getState(state string) *State {
	return f.states[state]
}

func (f *FSM) initClient(client Client) {
	curState := f.getState(client.State())
	if curState == nil {
		client.SetState(f.initialState)
		f.Emit(EvtTransition, NewBaseEvent(client, Undefined, f.initialState))
		curState = f.getState(client.State())
		if curState != nil {
			curState.runEvent(EvtEnter, f, client)
		}
	}
}

func (f *FSM) Transition(client Client, state string) {
	f.initClient(client)

	oldState := client.State()
	oldFsmState := f.getState(oldState)
	if oldFsmState == nil {
		return
	}

	// exit old state
	oldFsmState.runEvent(EvtExit, f, client)

	// set client new state
	client.SetState(state)

	f.Emit(EvtTransition, NewBaseEvent(client, oldState, state))

	newFsmState := f.getState(state)
	if newFsmState == nil {
		return
	}

	newFsmState.runEvent(EvtEnter, f, client)
}

func (f *FSM) DeferAndTransition(client Client, state string) {
	oldState := client.State()
	oldFsmState := f.getState(oldState)
	newFsmState := f.getState(state)

	f.Transition(client, state)

	if oldFsmState == nil {
		return
	}

	// 通过状态过滤无效事件
	if state != client.State() {
		return
	}

	f.runStateEvent(client, newFsmState, client.GetLastEvent(), client.GetLastPars()...)
}

func (f *FSM) runStateEvent(client Client, state *State, evtName string, pars ...interface{}) int {
	var ret int
	if state == nil {
		return ret
	}

	client.SetLastEvent(evtName)
	client.SetLastPars(pars)

	ret = state.runEvent(evtName, f, client, pars...)
	return ret
}

func (f *FSM) Handle(client Client, evtName string, pars ...interface{}) int {
	var ret int
	f.initClient(client)

	curState := f.getState(client.State())
	ret = f.runStateEvent(client, curState, evtName, pars...)

	return ret
}
