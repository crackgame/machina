package machina

import "reflect"

// State map[string]listener
type State map[string]interface{}

type States map[string]*State

// 优先执行指定的事件回调，如果没有就执行注册了"*"事件的函数
func (s State) runEvent(evtName string, fsm *FSM, client Client, pars ...interface{}) int {
	listener := s[evtName]
	if listener == nil {
		listener = s[AlwayEvent]
		if listener == nil {
			//fmt.Println("DEBUG:", client.State(), evtName, "func not found", s)
			return 0
		}
	}

	fn := reflect.ValueOf(listener)

	var values []reflect.Value

	// 先添加fsm和client做为前两个参数
	values = append(values, reflect.ValueOf(fsm))
	values = append(values, reflect.ValueOf(client))

	// 再添加自定义参数
	for i := 0; i < len(pars); i++ {
		if pars[i] == nil {
			values = append(values, reflect.New(fn.Type().In(i)).Elem())
		} else {
			values = append(values, reflect.ValueOf(pars[i]))
		}
	}

	//fmt.Println("DEBUG:", "before run event call fn,", evtName)
	rets := fn.Call(values)
	if len(rets) > 0 {
		return int(rets[0].Int())
	}

	return 0
}
