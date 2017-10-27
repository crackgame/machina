package main

import "github.com/crackgame/machina"
import "fmt"

type Unit struct {
	machina.BaseClient
}

func (u *Unit) unitState() string {
	return u.State()
}

func main() {
	testUnit := &Unit{}
	UnitFsm := machina.NewFSM(
		"idle",
		machina.States{
			// idle 状态处理逻辑
			"idle": {
				// _onEnter 事件触发回调函数
				"_onEnter": func(this *machina.FSM, unit *Unit) {
					fmt.Println("enter idle1")
					fmt.Println("enter idle2")
				},
				// _onExit 事件触发回调函数
				"_onExit": func(this *machina.FSM, unit *Unit) {
					fmt.Println("exit idle")
				},
				// tick 事件触发回调函数
				"tick": func(this *machina.FSM, unit *Unit, t int, dt int) {
					fmt.Println("idle tick1:", unit.unitState(), t, dt)
					this.DeferAndTransition(unit, "cd")
					fmt.Println("idle tick2:", unit.unitState(), t, dt)
				},
				// tick 事件触发回调函数
				"confirm": func(this *machina.FSM, unit *Unit, t int) {
					fmt.Println("idle confirm:", unit.unitState(), t)
				},
			},

			// cd 状态处理逻辑
			"cd": {
				// _onEnter 事件触发回调函数
				"_onEnter": func(this *machina.FSM, unit *Unit) {
					fmt.Println("enter cd")
				},
				// _onExit 事件触发回调函数
				"_onExit": func(this *machina.FSM, unit *Unit) {
					fmt.Println("exit cd")
				},
				// tick 事件触发回调函数
				"confirm": func(this *machina.FSM, unit *Unit, t int) {
					fmt.Println("cd confirm1:", unit.unitState(), t)
					this.DeferAndTransition(unit, "ready")
					fmt.Println("cd confirm2:,", unit.unitState(), t)
				},
			},

			// ready 状态处理逻辑
			"ready": {
				// * 事件触发回调函数
				"*": func(this *machina.FSM, unit *Unit) {
					fmt.Println("alway run by ready")
				},
				// _onEnter 事件触发回调函数
				"_onEnter": func(this *machina.FSM, unit *Unit) {
					fmt.Println("ready enter1:", unit.unitState())
					this.DeferAndTransition(unit, "idle")
					fmt.Println("ready enter2:", unit.unitState())
				},
				// _onExit 事件触发回调函数
				"_onExit": func(this *machina.FSM, unit *Unit) {
					fmt.Println("exit ready")
				},
			},
		},
	)

	UnitFsm.On("transition", func(event machina.Event) {
		unit := event.Client().(*Unit)
		fmt.Println("transition:", unit.unitState(), event.FromState(), "=>", event.ToState())
	})

	fmt.Println("1. UnitFsm.handle(unit, 'tick', 100, 5);")
	UnitFsm.Handle(testUnit, "tick", 100, 5)

	fmt.Println("2. UnitFsm.handle(unit, 'confirm', 100, 5);")
	UnitFsm.Handle(testUnit, "confirm", 100)
}
