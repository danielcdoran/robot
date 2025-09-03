package main

import (
	"testing"
)

func TestTurnRightStartNorth(t *testing.T) {
	sm := NewStateMachine(North)
	sm.SendEvent(TurnRight)
	result := sm.currentState
	expected := East
	if expected != result {
		t.Errorf("Expected Right turn result %+v different from result  position %+v", expected, result)
	}
}
func TestTurnRightStartEast(t *testing.T) {
	sm := NewStateMachine(East)
	sm.SendEvent(TurnRight)
	result := sm.currentState
	expected := South
	if expected != result {
		t.Errorf("Expected Right turn result %+v different from result  position %+v", expected, result)
	}
}
func TestTurnRightStartSouth(t *testing.T) {
	sm := NewStateMachine(South)
	sm.SendEvent(TurnRight)
	result := sm.currentState
	expected := West
	if expected != result {
		t.Errorf("Expected Right turn result %+v different from result  position %+v", expected, result)
	}
}
func TestTurnRightStartWest(t *testing.T) {
	sm := NewStateMachine(West)
	sm.SendEvent(TurnRight)
	result := sm.currentState
	expected := North
	if expected != result {
		t.Errorf("Expected Right turn result %+v different from result  position %+v", expected, result)
	}
}
