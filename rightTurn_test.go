package main

import (
	"testing"
)

func TestTurnRightStartNorth(t *testing.T) {
	pos := NewPosition(10, 10)
	sm := NewStateMachine(North, pos)
	posTopRight := NewPosition(49, 49)
	sm.SetTopRightCorner(posTopRight)
	sm.SendEvent(TurnRight)
	result := sm.currentState
	expected := East
	if expected != result {
		t.Errorf("Expected Right turn result %+v different from result  position %+v", expected, result)
	}
}
func TestTurnRightStartEast(t *testing.T) {
	pos := NewPosition(10, 10)
	sm := NewStateMachine(East, pos)
	posTopRight := NewPosition(49, 49)
	sm.SetTopRightCorner(posTopRight)
	sm.SendEvent(TurnRight)
	result := sm.currentState
	expected := South
	if expected != result {
		t.Errorf("Expected Right turn result %+v different from result  position %+v", expected, result)
	}
}
func TestTurnRightStartSouth(t *testing.T) {
	pos := NewPosition(10, 10)
	sm := NewStateMachine(South, pos)
	posTopRight := NewPosition(49, 49)
	sm.SetTopRightCorner(posTopRight)
	sm.SendEvent(TurnRight)
	result := sm.currentState
	expected := West
	if expected != result {
		t.Errorf("Expected Right turn result %+v different from result  position %+v", expected, result)
	}
}
func TestTurnRightStartWest(t *testing.T) {
	pos := NewPosition(10, 10)
	sm := NewStateMachine(West, pos)
	posTopRight := NewPosition(49, 49)
	sm.SetTopRightCorner(posTopRight)
	sm.SendEvent(TurnRight)
	result := sm.currentState
	expected := North
	if expected != result {
		t.Errorf("Expected Right turn result %+v different from result  position %+v", expected, result)
	}
}
