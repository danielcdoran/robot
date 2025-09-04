package main

import (
	"testing"
)

func TestTurnLeftStartNorth(t *testing.T) {
	pos := NewPosition(10, 10)
	sm := NewStateMachine(North, pos)
	posTopRight := NewPosition(49, 49)
	sm.SetTopRightCorner(posTopRight)
	sm.SendEvent(TurnLeft)
	result := sm.currentState
	expected := West
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}
func TestTurnLeftStartEast(t *testing.T) {
	pos := NewPosition(10, 10)
	sm := NewStateMachine(East, pos)
	posTopRight := NewPosition(49, 49)
	sm.SetTopRightCorner(posTopRight)
	sm.SendEvent(TurnLeft)
	result := sm.currentState
	expected := North
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}
func TestTurnLeftStartSouth(t *testing.T) {
	pos := NewPosition(10, 10)
	sm := NewStateMachine(South, pos)
	posTopRight := NewPosition(49, 49)
	sm.SetTopRightCorner(posTopRight)
	sm.SendEvent(TurnLeft)
	result := sm.currentState
	expected := East
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}
func TestTurnLeftStartWest(t *testing.T) {
	pos := NewPosition(10, 10)
	sm := NewStateMachine(West, pos)
	posTopRight := NewPosition(49, 49)
	sm.SetTopRightCorner(posTopRight)
	sm.SendEvent(TurnLeft)
	result := sm.currentState
	expected := South
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}
