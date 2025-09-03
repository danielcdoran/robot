package main

import "fmt"

type State int

const (
	North State = iota
	East
	South
)

type Event int

const (
	TurnLeft Event = iota
	TurnRight
	Move
)

type Action func()

type StateMachine struct {
	currentState State
	transitions  map[State]map[Event]State
	actions      map[State]map[Event]Action
}

func NewStateMachine(initialState State) *StateMachine {
	sm := &StateMachine{
		currentState: initialState,
		transitions:  make(map[State]map[Event]State),
		actions:      make(map[State]map[Event]Action),
	}

	sm.transitions[North] = map[Event]State{
		TurnLeft: East,
	}
	sm.transitions[East] = map[Event]State{
		TurnRight: South,
	}
	sm.transitions[South] = map[Event]State{
		Move: North,
	}

	sm.actions[North] = map[Event]Action{
		TurnLeft: func() { sm.currentState = South },
	}
	sm.actions[East] = map[Event]Action{
		TurnRight: func() { sm.currentState = North },
	}
	sm.actions[South] = map[Event]Action{
		Move: func() { sm.currentState = East },
	}

	return sm
}

func (sm *StateMachine) SendEvent(event Event) {
	if newState, ok := sm.transitions[sm.currentState][event]; ok {
		sm.currentState = newState
		if action, ok := sm.actions[sm.currentState][event]; ok {
			action()
		}
	} else {
		fmt.Println("Invalid transition")
	}
}

func main() {
	sm := NewStateMachine(North)
	fmt.Println(sm.currentState)
	sm.SendEvent(TurnLeft)
	fmt.Println(sm.currentState)
	sm.SendEvent(TurnRight)
	sm.SendEvent(Move)
}
