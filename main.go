package main

import (
	"fmt"
)

type State interface {
	Enter()
	Exit()
	Update(l *StateMachine)
}

type StateMachine struct {
	currentState State
	states       map[string]State
}

func NewStateMachine(initialState State) *StateMachine {
	sm := &StateMachine{
		currentState: initialState,
		states:       make(map[string]State),
	}

	sm.currentState.Enter()
	return sm
}

func (sm *StateMachine) setState(s State) {
	sm.currentState = s
	sm.currentState.Enter()
}

func (sm *StateMachine) Transition() {
	sm.currentState.Update(sm)
}

type RedLight struct{}

func (g RedLight) Enter() {
	fmt.Println("Red light is on. Stop driving.")

}
func (g RedLight) Exit() {}
func (g RedLight) Update(l *StateMachine) {
	l.setState(&GreenLight{})
}

type GreenLight struct{}

func (g GreenLight) Enter() {
	fmt.Println("Green light is on. You can drive.")

}
func (g GreenLight) Exit() {}
func (g GreenLight) Update(l *StateMachine) {
	l.setState(&YellowLight{})
}

type YellowLight struct{}

func (g YellowLight) Enter() {
	fmt.Println("Yellow light is on. Prepare to stop.")

}
func (g YellowLight) Exit() {}
func (g YellowLight) Update(l *StateMachine) {
	l.setState(&RedLight{})
}

type PurpleLight struct{}
func (g PurpleLight) Enter() {
	fmt.Println("Purple light is on. Prepare to stop.")

}
func (g PurpleLight) Exit() {}
func (g PurpleLight) Update(l *StateMachine) {
	l.setState(&RedLight{})
}

func main() {
	sm := NewStateMachine(&GreenLight{})

	for {
		sm.Transition()
	}
}
