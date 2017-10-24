package machina

type EventFunc func(*FSM, Client, ...interface{}) int

type State map[string]EventFunc

type States map[string]*State

func (s State) runEvent(evtName string, fsm *FSM, client Client, pars ...interface{}) int {
	fn := s[evtName]
	if fn == nil {
		//fmt.Println("DEBUG:", client.State(), evtName, "func not found", s)
		return 0
	}
	return fn(fsm, client, pars...)
}
