package main

import (
	"testing"
)

func TestMoveNorth(t *testing.T) {
	pos := NewPosition(10, 10)
	sm := NewStateMachine(North, pos)
	sm.SendEvent(Forward)
	result := sm.currentState
	expected := North
	if expected != result {
		t.Errorf("Expected direction %+v to be same but got different %+v", expected, result)
	}
	expectedPos := NewPosition(10, 11)
	if expectedPos != sm.position {
		t.Errorf("Expected positiond %+v to be same but got different %+v", expectedPos, sm.position)
	}
}

func TestMoveEast(t *testing.T) {
	pos := NewPosition(10, 10)
	sm := NewStateMachine(East, pos)
	sm.SendEvent(Forward)
	result := sm.currentState
	expected := East
	if expected != result {
		t.Errorf("Expected direction %+v to be same but got different %+v", expected, result)
	}
	expectedPos := NewPosition(11, 10)
	if expectedPos != sm.position {
		t.Errorf("Expected positiond %+v to be same but got different %+v", expectedPos, sm.position)
	}
}
func TestMoveSouth(t *testing.T) {
	pos := NewPosition(10, 10)
	sm := NewStateMachine(South, pos)
	sm.SendEvent(Forward)
	result := sm.currentState
	expected := South
	if expected != result {
		t.Errorf("Expected direction %+v to be same but got different %+v", expected, result)
	}
	expectedPos := NewPosition(10, 9)
	if expectedPos != sm.position {
		t.Errorf("Expected positiond %+v to be same but got different %+v", expectedPos, sm.position)
	}
}
func TestMoveWest(t *testing.T) {
	pos := NewPosition(10, 10)
	sm := NewStateMachine(West, pos)
	sm.SendEvent(Forward)
	result := sm.currentState
	expected := West
	if expected != result {
		t.Errorf("Expected direction %+v to be same but got different %+v", expected, result)
	}
	expectedPos := NewPosition(9, 10)
	if expectedPos != sm.position {
		t.Errorf("Expected positiond %+v to be same but got different %+v", expectedPos, sm.position)
	}
}
