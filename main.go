package main

import (
	"fmt"
)

type State interface {
	MoveInDirection(l *StateMachine)
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
func (p RobotDetails) String() string {
	return fmt.Sprintf("%v (%v years)", p.position, p.facing)
}

type StateMachine struct {
	currentState     State
	states           map[string]State
	currentStateData StateData
	planet           [50][50]bool
}

func NewStateMachine(initialState RobotDetails) *StateMachine {
	sm := &StateMachine{
		currentState: &GreenLight{},
		states:       make(map[string]State),
	}
	sm.currentStateData = NewStateData(Open, initialState)
	// sm.currentState.MoveInDirection(l * StateMachine)
	return sm
}
func (p StateMachine) String() string {
	return fmt.Sprintf("%v", p.currentStateData)
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

func moveInDirection(detail RobotDetails) RobotDetails {
	val := NewRobotDetails(detail.position, detail.facing)
	xadd := []int64{0, 1, 0, -1}
	yadd := []int64{1, 0, -1, 0}
	val.position.xpos = val.position.xpos + xadd[val.facing]
	val.position.ypos = val.position.ypos + yadd[val.facing]
	return val
}
func (sm *StateMachine) setState(s State) {
	sm.currentState = s
	// sm.currentState.MoveInDirection(l.currentStateData.position)
}

func (sm *StateMachine) TurnLeft() {
	sm.currentState.TurnLeft(sm)
}
func (sm *StateMachine) TurnRight() {
	sm.currentState.TurnRight(sm)
}
func (sm *StateMachine) MoveInDirection() {
	sm.currentState.MoveInDirection(sm)
}

type RedLight struct{}

func (g RedLight) MoveInDirection(l *StateMachine) {
	fmt.Println("Red light is on. Stop driving.")
	l.currentStateData.position = moveInDirection(l.currentStateData.position)
}
func (g RedLight) TurnRight(l *StateMachine) {
	l.currentStateData.position = turnRight(l.currentStateData.position)
}
func (g RedLight) TurnLeft(l *StateMachine) {
	l.setState(&GreenLight{})
	l.currentStateData.position = turnLeft(l.currentStateData.position)
}

type GreenLight struct{}

func (g GreenLight) MoveInDirection(l *StateMachine) {
	fmt.Println("Green light is on. You can drive.")
	l.currentStateData.position = moveInDirection(l.currentStateData.position)
}
func (g GreenLight) TurnRight(l *StateMachine) {
	l.currentStateData.position = turnRight(l.currentStateData.position)
}
func (g GreenLight) TurnLeft(l *StateMachine) {
	l.setState(&YellowLight{})
	l.currentStateData.position = turnLeft(l.currentStateData.position)
}

type YellowLight struct{}

func (g YellowLight) MoveInDirection(l *StateMachine) {
	fmt.Println("Yellow light is on. Prepare to stop.")
	l.currentStateData.position = moveInDirection(l.currentStateData.position)
}
func (g YellowLight) TurnRight(l *StateMachine) {
	l.currentStateData.position = turnRight(l.currentStateData.position)
}
func (g YellowLight) TurnLeft(l *StateMachine) {
	l.setState(&RedLight{})
	l.currentStateData.position = turnLeft(l.currentStateData.position)
}

type PurpleLight struct{}

func (g PurpleLight) MoveInDirection(l *StateMachine) {
	fmt.Println("Purple light is on. Prepare to stop.")
	l.currentStateData.position = moveInDirection(l.currentStateData.position)
}
func (g PurpleLight) TurnRight(l *StateMachine) {
	l.currentStateData.position = turnRight(l.currentStateData.position)
}
func (g PurpleLight) TurnLeft(l *StateMachine) {
	l.setState(&RedLight{})
	l.currentStateData.position = turnLeft(l.currentStateData.position)
}

func main() {
	pos := NewRobotPosition(10, 10)
	det := NewRobotDetails(pos, North)
	sm := NewStateMachine(det)

	sm.TurnLeft()
	fmt.Println(sm)
	sm.TurnRight()
	fmt.Println(sm)
	sm.MoveInDirection()
	fmt.Println(sm)
}
