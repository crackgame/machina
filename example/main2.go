package main

import "github.com/crackgame/machina"
import "fmt"

type Unit struct {
	machina.BaseClient

	k int
}

func (u *Unit) unitState() string {
	return u.State()
}

func main() {
	testUnit := &Unit{k: 1}
	UnitFsm := machina.NewFSM(
		"idle",
		machina.States{
			// idle 状态处理逻辑
			"idle": {
				// _onEnter 事件触发回调函数
				"_onEnter": func(this *machina.FSM, client machina.Client, pars ...interface{}) int {
					unit := client.(*Unit)
					fmt.Println("enter idle1")
					if unit.k == 2 {
						fmt.Println("11. UnitFsm.deferAndTransition(unit, 'cd');")
						this.DeferAndTransition(unit, "cd")
					}
					unit.k = 2
					fmt.Println("enter idle2")
					return 0
				},
				// _onExit 事件触发回调函数
				"_onExit": func(this *machina.FSM, client machina.Client, pars ...interface{}) int {
					fmt.Println("exit idle")
					return 0
				},
				// tick 事件触发回调函数
				"tick": func(this *machina.FSM, client machina.Client, pars ...interface{}) int {
					unit := client.(*Unit)
					t := pars[0].(int)
					dt := pars[1].(int)

					fmt.Println("idle tick1:", unit.unitState(), t, dt)
					this.DeferAndTransition(unit, "cd")
					fmt.Println("idle tick2:", unit.unitState(), t, dt)
					return 0
				},
				// tick 事件触发回调函数
				"confirm": func(this *machina.FSM, client machina.Client, pars ...interface{}) int {
					unit := client.(*Unit)
					t := pars[0].(int)

					fmt.Println("idle confirm:", unit.unitState(), t)
					return 0
				},
			},

			// cd 状态处理逻辑
			"cd": {
				// * 事件触发回调函数
				"*": func(this *machina.FSM, client machina.Client, pars ...interface{}) int {
					fmt.Println("alway run by cd")
					return 0
				},
				// _onEnter 事件触发回调函数
				"_onEnter": func(this *machina.FSM, client machina.Client, pars ...interface{}) int {
					fmt.Println("enter cd")
					return 0
				},
				// _onExit 事件触发回调函数
				"_onExit": func(this *machina.FSM, client machina.Client, pars ...interface{}) int {
					fmt.Println("exit cd")
					return 0
				},
				// tick 事件触发回调函数
				"confirm": func(this *machina.FSM, client machina.Client, pars ...interface{}) int {
					unit := client.(*Unit)
					t := pars[0].(int)

					fmt.Println("cd confirm1:", unit.unitState(), t)
					this.DeferAndTransition(unit, "ready")
					fmt.Println("cd confirm2:,", unit.unitState(), t)
					return 0
				},
			},

			// ready 状态处理逻辑
			"ready": {
				// * 事件触发回调函数
				"*": func(this *machina.FSM, client machina.Client, pars ...interface{}) int {
					fmt.Println("alway run by ready")
					return 0
				},
				// _onEnter 事件触发回调函数
				"_onEnter": func(this *machina.FSM, client machina.Client, pars ...interface{}) int {
					fmt.Println("enter ready")
					return 0
				},
				// _onExit 事件触发回调函数
				"_onExit": func(this *machina.FSM, client machina.Client, pars ...interface{}) int {
					// TODO
					fmt.Println("exit ready")
					return 0
				},
				// confirm 事件触发回调函数
				"confirm": func(this *machina.FSM, client machina.Client, pars ...interface{}) int {
					// TODO
					unit := client.(*Unit)
					t := pars[0].(int)

					fmt.Println("ready confirm:", unit.unitState(), t)
					return 0
				},
				// tick 事件触发回调函数
				"tick": func(this *machina.FSM, client machina.Client, pars ...interface{}) int {
					// TODO
					unit := client.(*Unit)
					t := pars[0].(int)
					dt := pars[1].(int)

					fmt.Println("ready tick1:", unit.unitState(), t, dt)
					this.DeferAndTransition(unit, "idle")
					fmt.Println("ready tick2:", unit.unitState(), t, dt)
					return 0
				},
			},
		},
	)

	UnitFsm.On("transition", func(event machina.Event) {
		unit := event.Client().(*Unit)
		fmt.Println("transition:", unit.unitState(), event.FromState(), "=>", event.ToState())
	})

	fmt.Println("1. UnitFsm.handle(unit, 'confirm', 100, 5);")
	UnitFsm.Handle(testUnit, "confirm", 100, 5)
	//debug(unit)

	fmt.Println("2. UnitFsm.deferAndTransition(unit, 'cd');")
	UnitFsm.DeferAndTransition(testUnit, "cd")
	//debug(unit)

	fmt.Println("3. UnitFsm.handle(unit, 'tick', 100, 5);")
	UnitFsm.Handle(testUnit, "tick", 100, 5)
}
