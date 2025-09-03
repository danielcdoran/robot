package main

import (
	"testing"
)

func TestMoveXLessThan0(t *testing.T) {
	pos := NewPosition(0, 0)
	sm := NewStateMachine(West, pos)
	sm.SendEvent(Move)
	result := sm.currentState
	expected := West
	if expected != result {
		t.Errorf("Expected direction %+v to be same but got different %+v", expected, result)
	}
	expectedPos := NewPosition(0, 0)
	if expectedPos != sm.position {
		t.Errorf("Expected positiond %+v to be same but got different %+v", expectedPos, sm.position)
	}
	if !sm.planetScent[0][0] {
		t.Errorf("Expected scented (0,0) to be true but got %+v", sm.planetScent[0][0])
	}
}

func TestMoveYLessThan0(t *testing.T) {
	pos := NewPosition(49, 0)
	sm := NewStateMachine(South, pos)
	sm.SendEvent(Move)
	result := sm.currentState
	expected := South
	if expected != result {
		t.Errorf("Expected direction %+v to be same but got different %+v", expected, result)
	}
	expectedPos := NewPosition(49, 0)
	if expectedPos != sm.position {
		t.Errorf("Expected position %+v to be same but got different %+v", expectedPos, sm.position)
	}
	if !sm.planetScent[49][0] {
		t.Errorf("Expected scented (49,0) to be true but got %+v", sm.planetScent[49][0])
	}
}

func TestMoveYGTThan50(t *testing.T) {
	pos := NewPosition(49, 49)
	sm := NewStateMachine(North, pos)
	sm.SendEvent(Move)
	result := sm.currentState
	expected := North
	if expected != result {
		t.Errorf("Expected direction %+v to be same but got different %+v", expected, result)
	}
	expectedPos := NewPosition(49, 49)
	if expectedPos != sm.position {
		t.Errorf("Expected position %+v to be same but got different %+v", expectedPos, sm.position)
	}
	if !sm.planetScent[49][49] {
		t.Errorf("Expected scented (49,0) to be true but got %+v", sm.planetScent[49][49])
	}
}

func TestMoveXGTThan50(t *testing.T) {
	pos := NewPosition(49, 49)
	sm := NewStateMachine(East, pos)
	sm.SendEvent(Move)
	result := sm.currentState
	expected := East
	if expected != result {
		t.Errorf("Expected direction %+v to be same but got different %+v", expected, result)
	}
	expectedPos := NewPosition(49, 49)
	if expectedPos != sm.position {
		t.Errorf("Expected position %+v to be same but got different %+v", expectedPos, sm.position)
	}
	if !sm.planetScent[49][49] {
		t.Errorf("Expected scented (49,0) to be true but got %+v", sm.planetScent[49][49])
	}
}
