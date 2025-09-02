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
func (p StateData) OutputRobotPosition() string {
	return p.position.OutputRobotPosition()
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
	return fmt.Sprintf("%v %v", p.position, p.facing)
}

func (p RobotDetails) OutputRobotPosition() string {
	var direction string
	switch p.facing {

	case 0:
		direction = "N"
	case 1:
		direction = "E"
	case 2:
		direction = "S"
	case 3:
		direction = "W"
	default:
		fmt.Println("Direction %s not in facing list", p.facing)
	}
	return fmt.Sprintf("%d %d %s", p.position.xpos, p.position.ypos, direction)
}

type StateMachine struct {
	currentState     State
	states           map[string]State
	currentStateData StateData
	planetScent      [50][50]bool // false - no scent
}

func NewStateMachine(initialState RobotDetails) *StateMachine {
	sm := &StateMachine{
		currentState: &Action{},
		states:       make(map[string]State),
	}
	sm.currentStateData = NewStateData(Open, initialState)
	return sm
}
func (p StateMachine) String() string {
	return fmt.Sprintf("%v", p.currentStateData)
}

func (p StateMachine) OutputRobotPosition() string {
	return p.currentStateData.OutputRobotPosition()
}
func (p *StateMachine) hasScent() bool {
	val := p.planetScent[p.currentStateData.position.position.xpos][p.currentStateData.position.position.ypos]
	return val
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

// If position is scented ie  true then no movement is allowed
// If new position is outside area then no movement and setto scented
func moveInDirection(detail RobotDetails, outsideArea *bool) RobotDetails {
	val := NewRobotDetails(detail.position, detail.facing)
	xadd := []int64{0, 1, 0, -1}
	yadd := []int64{1, 0, -1, 0}
	newxpos := val.position.xpos + xadd[val.facing]
	newypos := val.position.ypos + yadd[val.facing]
	if isOutsideArea(newxpos, newypos) {
		*outsideArea = true
		return val
	}
	*outsideArea = false
	val.position.xpos = newxpos
	val.position.ypos = newypos
	return val
}

func isOutsideArea(xpos int64, ypos int64) bool {
	if xpos < 0 {
		return true
	}
	if xpos > 50 {

		return true
	}
	if ypos < 0 {
		return true
	}
	if ypos > 50 {
		return true
	}
	return false
}

// func (sm *StateMachine) setState(s State) {
// 	sm.currentState = s
// }

func (sm *StateMachine) TurnLeft() {
	sm.currentState.TurnLeft(sm)
}
func (sm *StateMachine) TurnRight() {
	sm.currentState.TurnRight(sm)
}
func (sm *StateMachine) MoveInDirection() {
	if sm.hasScent() {
		return // no movement allowed
	}
	sm.currentState.MoveInDirection(sm)
}

type Action struct{}

func (g Action) MoveInDirection(l *StateMachine) {
	var outSideArea bool
	l.currentStateData.position = moveInDirection(l.currentStateData.position, &outSideArea)
	if outSideArea {
		l.planetScent[l.currentStateData.position.position.xpos][l.currentStateData.position.position.ypos] = true
	}
}
func (g Action) TurnRight(l *StateMachine) {
	l.currentStateData.position = turnRight(l.currentStateData.position)
}
func (g Action) TurnLeft(l *StateMachine) {
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
	fmt.Println()
	fmt.Println(sm.OutputRobotPosition())
}
