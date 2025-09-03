package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	Forward
)

type Action func()
type Scented [50][50]bool

type StateMachine struct {
	currentState State
	transitions  map[State]map[Event]State
	actions      map[State]map[Event]Action
	position     Position
	planetScent  Scented // false - no scent
}

func NewStateMachine(initialState State, pos Position) *StateMachine {
	sm := NewStateMachineSetStart()
	sm.currentState = initialState
	sm.position = pos
	return sm
}

func NewStateMachineSetStart() *StateMachine {
	sm := &StateMachine{
		transitions: make(map[State]map[Event]State),
		actions:     make(map[State]map[Event]Action),
	}

	sm.transitions[North] = map[Event]State{
		TurnLeft:  West,
		TurnRight: East,
		Forward:   North,
	}
	sm.transitions[East] = map[Event]State{
		TurnLeft:  North,
		TurnRight: South,
		Forward:   East,
	}
	sm.transitions[South] = map[Event]State{
		TurnLeft:  East,
		TurnRight: West,
		Forward:   South,
	}
	sm.transitions[West] = map[Event]State{
		TurnLeft:  South,
		TurnRight: North,
		Forward:   West,
	}

	sm.actions[North] = map[Event]Action{
		TurnLeft:  func() { /* sm.currentState = West */ },
		TurnRight: func() { /* sm.currentState = West */ },
		Forward: func() {
			newpos := NewPositionCopy(sm.position)
			newpos.ypos = newpos.ypos + 1
			sm.position = checkScented(&sm.planetScent, sm.position, newpos)
		},
	}
	sm.actions[East] = map[Event]Action{
		TurnLeft:  func() { /* sm.currentState = North */ },
		TurnRight: func() { /* sm.currentState = West */ },
		Forward: func() {
			newpos := NewPositionCopy(sm.position)
			newpos.xpos = newpos.xpos + 1
			sm.position = checkScented(&sm.planetScent, sm.position, newpos)
		},
	}
	sm.actions[South] = map[Event]Action{
		TurnLeft:  func() { /* sm.currentState = East */ },
		TurnRight: func() { /* sm.currentState = West */ },
		Forward: func() {
			newpos := NewPositionCopy(sm.position)
			newpos.ypos = newpos.ypos - 1
			sm.position = checkScented(&sm.planetScent, sm.position, newpos)
		},
	}
	sm.actions[West] = map[Event]Action{
		TurnLeft:  func() { /* sm.currentState = South */ },
		TurnRight: func() { /* sm.currentState = West */ },
		Forward: func() {
			newpos := NewPositionCopy(sm.position)
			newpos.xpos = newpos.xpos - 1
			sm.position = checkScented(&sm.planetScent, sm.position, newpos)
		},
	}

	return sm
}

// func NewStateMachine(initialState State, pos Position) *StateMachine {
// 	sm := &StateMachine{
// 		currentState: initialState,
// 		transitions:  make(map[State]map[Event]State),
// 		actions:      make(map[State]map[Event]Action),
// 		position:     pos,
// 	}

// 	sm.transitions[North] = map[Event]State{
// 		TurnLeft:  West,
// 		TurnRight: East,
// 		Forward:   North,
// 	}
// 	sm.transitions[East] = map[Event]State{
// 		TurnLeft:  North,
// 		TurnRight: South,
// 		Forward:   East,
// 	}
// 	sm.transitions[South] = map[Event]State{
// 		TurnLeft:  East,
// 		TurnRight: West,
// 		Forward:   South,
// 	}
// 	sm.transitions[West] = map[Event]State{
// 		TurnLeft:  South,
// 		TurnRight: North,
// 		Forward:   West,
// 	}

// 	sm.actions[North] = map[Event]Action{
// 		TurnLeft:  func() { /* sm.currentState = West */ },
// 		TurnRight: func() { /* sm.currentState = West */ },
// 		Forward: func() {
// 			newpos := NewPositionCopy(sm.position)
// 			newpos.ypos = newpos.ypos + 1
// 			sm.position = checkScented(&sm.planetScent, sm.position, newpos)
// 		},
// 	}
// 	sm.actions[East] = map[Event]Action{
// 		TurnLeft:  func() { /* sm.currentState = North */ },
// 		TurnRight: func() { /* sm.currentState = West */ },
// 		Forward: func() {
// 			newpos := NewPositionCopy(sm.position)
// 			newpos.xpos = newpos.xpos + 1
// 			sm.position = checkScented(&sm.planetScent, sm.position, newpos)
// 		},
// 	}
// 	sm.actions[South] = map[Event]Action{
// 		TurnLeft:  func() { /* sm.currentState = East */ },
// 		TurnRight: func() { /* sm.currentState = West */ },
// 		Forward: func() {
// 			newpos := NewPositionCopy(sm.position)
// 			newpos.ypos = newpos.ypos - 1
// 			sm.position = checkScented(&sm.planetScent, sm.position, newpos)
// 		},
// 	}
// 	sm.actions[West] = map[Event]Action{
// 		TurnLeft:  func() { /* sm.currentState = South */ },
// 		TurnRight: func() { /* sm.currentState = West */ },
// 		Forward: func() {
// 			newpos := NewPositionCopy(sm.position)
// 			newpos.xpos = newpos.xpos - 1
// 			sm.position = checkScented(&sm.planetScent, sm.position, newpos)
// 		},
// 	}

//		return sm
//	}
func (sm *StateMachine) InitialPosition(state State, pos Position) {
	sm.currentState = state
	sm.position = pos
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

// Copy contructor
func NewPositionCopy(pos Position) Position {
	newPos := new(Position)
	newPos.xpos = pos.xpos
	newPos.ypos = pos.ypos
	return *newPos
}

func isOutsideArea(pos Position) bool {
	xpos := pos.xpos
	ypos := pos.ypos
	if xpos <= 0 {
		return true
	}
	if xpos >= 50 {

		return true
	}
	if ypos <= 0 {
		return true
	}
	if ypos >= 50 {
		return true
	}
	return false
}
func checkScented(scented *Scented, pos Position, movedPos Position) Position {
	if scented[pos.xpos][pos.ypos] {
		return pos
	}
	if isOutsideArea(movedPos) {
		scented[pos.xpos][pos.ypos] = true
		return pos
	}
	return movedPos
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getPosition(input string) (Position, State) {
	parts := strings.Split(input, " ")
	xpos, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	ypos, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	pos := NewPosition(xpos, ypos)

	state := letterToState(parts[2])
	return pos, state
}

func letterToState(input string) State {
	switch input {
	case "N":
		return 0
	case "E":
		return 1
	case "s":
		return 2
	case "W":
		return 3
	}
	return 0
}

func runCommands(input []string) {
	sm := NewStateMachineSetStart()
	for i := 1; i < len(input); i = i + 2 {
		pos, state := getPosition(input[i])
		sm.InitialPosition(state, pos)
		commands := input[i+1]
		for j := 0; j < len(commands); j++ {

			switch commands[j] {
			case 'L':
				sm.SendEvent(TurnLeft)
			case 'R':
				sm.SendEvent((TurnRight))
			case 'F':
				sm.SendEvent(Forward)
			}
		}
		fmt.Println("Commands  %s produce position  %s", commands, sm.String())
	}

}
func main() {
	// pos := NewPosition(10, 10)
	// sm := NewStateMachine(North, pos)
	// fmt.Println(sm)
	// sm.SendEvent(TurnRight)
	// fmt.Println(sm)
	// sm.SendEvent(TurnRight)
	// sm.SendEvent(Forward)
	// fmt.Println(sm)
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	runCommands(lines)

}
