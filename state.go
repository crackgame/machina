package machina

type EventFunc func(*FSM, Client, ...interface{}) int

type State map[string]EventFunc

type States map[string]*State

// 优先执行指定的事件回调，如果没有就执行注册了"*"事件的函数
func (s State) runEvent(evtName string, fsm *FSM, client Client, pars ...interface{}) int {
	fn := s[evtName]
	if fn == nil {
		fn = s[AlwayEvent]
		if fn == nil {
			//fmt.Println("DEBUG:", client.State(), evtName, "func not found", s)
			return 0
		}
	}
	return fn(fsm, client, pars...)
}
