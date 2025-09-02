package main

import (
	"fmt"
)

type State interface {
	Enter()
	Exit()
	Update(l *StateMachine)
}

type StateRobot int
const (
	Open StateRobot = iota
	InProgress
	Closed
)

type StateData struct {
	state    StateRobot
	position RobotDetails
}

func NewStateData(state StateRobot, pos RobotDetails) StateData {
	object := new(StateData)
	object.state = state
	object.position = pos
	return *object
}

type RobotPosition struct {
	xpos int64
	ypos int64
}

func NewRobotPosition(xpos int64, ypos int64) RobotPosition {
	pos := new(RobotPosition)
	pos.xpos = xpos
	pos.ypos = ypos
	return *pos
}

type Direction int

const ( // note : directions are clockwise
	North Direction = iota
	East
	South
	West
)


type RobotDetails struct {
	position RobotPosition
	facing   Direction
}

func NewRobotDetails(position RobotPosition, facing Direction) RobotDetails {
	pos := new(RobotDetails)
	pos.position = position
	pos.facing = facing
	return *pos
}

type StateMachine struct {
	currentState State
	states       map[string]State
		currentStateData StateData
	planet           [50][50]bool
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
