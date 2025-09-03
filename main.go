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
	position     Position
}

func NewStateMachine(initialState State, pos Position) *StateMachine {
	sm := &StateMachine{
		currentState: initialState,
		transitions:  make(map[State]map[Event]State),
		actions:      make(map[State]map[Event]Action),
		position:     pos,
	}

	sm.transitions[North] = map[Event]State{
		TurnLeft:  West,
		TurnRight: East,
		Move:      North,
	}
	sm.transitions[East] = map[Event]State{
		TurnLeft:  North,
		TurnRight: South,
		Move:      East,
	}
	sm.transitions[South] = map[Event]State{
		TurnLeft:  East,
		TurnRight: West,
		Move:      South,
	}
	sm.transitions[West] = map[Event]State{
		TurnLeft:  South,
		TurnRight: North,
		Move:      West,
	}

	sm.actions[North] = map[Event]Action{
		TurnLeft:  func() { /* sm.currentState = West */ },
		TurnRight: func() { /* sm.currentState = West */ },
		Move:      func() { sm.position.ypos = sm.position.ypos + 1 },
	}
	sm.actions[East] = map[Event]Action{
		TurnLeft:  func() { /* sm.currentState = North */ },
		TurnRight: func() { /* sm.currentState = West */ },
		Move:      func() { sm.position.xpos = sm.position.xpos + 1 },
	}
	sm.actions[South] = map[Event]Action{
		TurnLeft:  func() { /* sm.currentState = East */ },
		TurnRight: func() { /* sm.currentState = West */ },
		Move:      func() { sm.position.ypos = sm.position.ypos - 1 },
	}
	sm.actions[West] = map[Event]Action{
		TurnLeft:  func() { /* sm.currentState = South */ },
		TurnRight: func() { /* sm.currentState = West */ },
		Move:      func() { sm.position.xpos = sm.position.xpos - 1 },
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
func (p StateMachine) String() string {
	return fmt.Sprintf("Facing %v  (%v,%v)", p.currentState, p.position.xpos, p.position.ypos)
}

type Position struct {
	xpos int64
	ypos int64
}

func NewPosition(xpos int64, ypos int64) Position {
	pos := new(Position)
	pos.xpos = xpos
	pos.ypos = ypos
	return *pos
}

func main() {
	pos := NewPosition(10, 10)
	sm := NewStateMachine(North, pos)
	fmt.Println(sm)
	sm.SendEvent(TurnRight)
	fmt.Println(sm)
	sm.SendEvent(TurnRight)
	sm.SendEvent(Move)
	fmt.Println(sm)
}
