package main

import (
	"fmt"
)

type State interface {
	Enter()
	TurnRight(l *StateMachine)
	TurnLeft(l *StateMachine)
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
	currentState     State
	states           map[string]State
	currentStateData StateData
	planet           [50][50]bool
}

func NewStateMachine(initialState State) *StateMachine {
	sm := &StateMachine{
		currentState: initialState,
		states:       make(map[string]State),
	}
	pos := NewRobotPosition(10, 10)
	det := NewRobotDetails(pos, North)
	sm.currentStateData = NewStateData(Open, det)
	sm.currentState.Enter()
	return sm
}
func NewStateMachine2(initialState RobotDetails) *StateMachine {
	sm := &StateMachine{
		currentState: &GreenLight{},
		states:       make(map[string]State),
	}
	pos := NewRobotPosition(10, 10)
	det := NewRobotDetails(pos, North)
	sm.currentStateData = NewStateData(Open, det)
	sm.currentState.Enter()
	return sm
}
func turnLeft(detail RobotDetails) RobotDetails {
	val := NewRobotDetails(detail.position, detail.facing)
	newDirection := []Direction{West, North, East, South}
	val.facing = newDirection[val.facing]
	return val
}

func turnRight(detail RobotDetails) RobotDetails {
	val := NewRobotDetails(detail.position, detail.facing)
	newDirection := []Direction{East, South, West, North}
	val.facing = newDirection[val.facing]
	return val
}
func (sm *StateMachine) setState(s State) {
	sm.currentState = s
	sm.currentState.Enter()
}

func (sm *StateMachine) Transition() {
	sm.currentState.TurnLeft(sm)
}

type RedLight struct{}

func (g RedLight) Enter() {
	fmt.Println("Red light is on. Stop driving.")

}
func (g RedLight) TurnRight(l *StateMachine) {
	l.currentStateData.position = turnRight(l.currentStateData.position)
}
func (g RedLight) TurnLeft(l *StateMachine) {
	l.setState(&GreenLight{})
	l.currentStateData.position = turnLeft(l.currentStateData.position)
}

type GreenLight struct{}

func (g GreenLight) Enter() {
	fmt.Println("Green light is on. You can drive.")

}
func (g GreenLight) TurnRight(l *StateMachine) {
	l.currentStateData.position = turnRight(l.currentStateData.position)
}
func (g GreenLight) TurnLeft(l *StateMachine) {
	l.setState(&YellowLight{})
	l.currentStateData.position = turnLeft(l.currentStateData.position)
}

type YellowLight struct{}

func (g YellowLight) Enter() {
	fmt.Println("Yellow light is on. Prepare to stop.")
}
func (g YellowLight) TurnRight(l *StateMachine) {
	l.currentStateData.position = turnRight(l.currentStateData.position)
}
func (g YellowLight) TurnLeft(l *StateMachine) {
	l.setState(&RedLight{})
	l.currentStateData.position = turnLeft(l.currentStateData.position)
}

type PurpleLight struct{}

func (g PurpleLight) Enter() {
	fmt.Println("Purple light is on. Prepare to stop.")

}
func (g PurpleLight) TurnRight(l *StateMachine) {
	l.currentStateData.position = turnRight(l.currentStateData.position)
}
func (g PurpleLight) TurnLeft(l *StateMachine) {
	l.setState(&RedLight{})
	l.currentStateData.position = turnLeft(l.currentStateData.position)
}

func main() {
	sm := NewStateMachine(&GreenLight{})
	sm.Transition()
	// for {
	// 	sm.Transition()
	// }
}
