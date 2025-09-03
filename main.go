package main

import "fmt"

type State int

const (
	North State = iota
	East
	South
	West
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
		TurnLeft:  West,
		TurnRight: West,
	}
	sm.transitions[East] = map[Event]State{
		TurnLeft:  North,
		TurnRight: West,
	}
	sm.transitions[South] = map[Event]State{
		TurnLeft:  East,
		TurnRight: West,
	}
	sm.transitions[West] = map[Event]State{
		TurnLeft:  South,
		TurnRight: West,
	}
	// sm.transitions[East] = map[Event]State{
	// 	TurnRight: South,
	// }
	// sm.transitions[South] = map[Event]State{
	// 	Move: North,
	// }

	sm.actions[North] = map[Event]Action{
		TurnLeft:  func() { sm.currentState = West },
		TurnRight: func() { sm.currentState = West },
	}
	sm.actions[East] = map[Event]Action{
		TurnLeft:  func() { sm.currentState = North },
		TurnRight: func() { sm.currentState = West },
	}
	sm.actions[South] = map[Event]Action{
		TurnLeft:  func() { sm.currentState = East },
		TurnRight: func() { sm.currentState = West },
	}
	sm.actions[West] = map[Event]Action{
		TurnLeft:  func() { sm.currentState = South },
		TurnRight: func() { sm.currentState = West },
	}

	sm.actions[East] = map[Event]Action{
		TurnLeft:  func() { sm.currentState = North },
		TurnRight: func() { sm.currentState = North },
	}
	sm.actions[South] = map[Event]Action{
		TurnLeft:  func() { sm.currentState = East },
		TurnRight: func() { sm.currentState = North },
	}
	sm.actions[North] = map[Event]Action{
		TurnLeft:  func() { sm.currentState = North },
		TurnRight: func() { sm.currentState = North },
	}
	sm.actions[West] = map[Event]Action{
		TurnLeft:  func() { sm.currentState = East },
		TurnRight: func() { sm.currentState = North },
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
	sm.SendEvent(TurnLeft)
	// sm.SendEvent(Move)
}
