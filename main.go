package main

import "fmt"

type State int

const (
	Open State = iota
	InProgress
	Closed
)

type Event int

const (
	StartProgress Event = iota
	Close
	Reopen
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

	sm.transitions[Open] = map[Event]State{
		StartProgress: InProgress,
	}
	sm.transitions[InProgress] = map[Event]State{
		Close: Closed,
	}
	sm.transitions[Closed] = map[Event]State{
		Reopen: Open,
	}

	sm.actions[Open] = map[Event]Action{
		StartProgress: func() { fmt.Println("Ticket is now in progress") },
	}
	sm.actions[InProgress] = map[Event]Action{
		Close: func() { fmt.Println("Ticket is now closed") },
	}
	sm.actions[Closed] = map[Event]Action{
		Reopen: func() { fmt.Println("Ticket is reopened") },
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
	sm := NewStateMachine(Open)
	sm.SendEvent(StartProgress)
	sm.SendEvent(Close)
	sm.SendEvent(Reopen)
}
