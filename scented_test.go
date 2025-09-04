package main

import (
	"testing"
)

//  scented = yes outtside are dont movedPos			TestScentedMovedIsOutsideArea
//  scented = yes  inside area movedPos
//  scented = no outside area set scented make movedPos   TestOutsideAreaSetScented
//  scented = no inside area move

func TestOutsideAreaSetScented(t *testing.T) {
	posTopRight := NewPosition(49, 49)
	pos := NewPosition(0, 0)
	sm := NewStateMachine(West, pos) // new so Scented = false
	sm.SetTopRightCorner(posTopRight)
	sm.SendEvent(Forward)
	result := sm.currentState
	expected := West
	if expected != result {
		t.Errorf("Expected direction %+v to be same but got different %+v", expected, result)
	}
	expectedPos := NewPosition(-1, 0)
	if expectedPos != sm.position {
		t.Errorf("Expected positiond %+v to be same but got different %+v", expectedPos, sm.position)
	}
	if !sm.planetScent[0][0] {
		t.Errorf("Expected scented (0,0) to be true but got %+v", sm.planetScent[0][0])
	}
}

func TestScentedMovedIsOutsideArea(t *testing.T) {
	posTopRight := NewPosition(49, 49)
	pos := NewPosition(0, 0)
	sm := NewStateMachine(West, pos)
	sm.planetScent[0][0] = true
	sm.SetTopRightCorner(posTopRight)
	sm.SendEvent(Forward)
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

func TestScentedMovedIsInsideArea(t *testing.T) {
	posTopRight := NewPosition(49, 49)
	pos := NewPosition(0, 0)
	sm := NewStateMachine(East, pos)
	sm.planetScent[0][0] = true
	sm.SetTopRightCorner(posTopRight)
	sm.SendEvent(Forward)
	result := sm.currentState
	expected := East
	if expected != result {
		t.Errorf("Expected direction %+v to be same but got different %+v", expected, result)
	}
	expectedPos := NewPosition(1, 0)
	if expectedPos != sm.position {
		t.Errorf("Expected positiond %+v to be same but got different %+v", expectedPos, sm.position)
	}
	if !sm.planetScent[0][0] {
		t.Errorf("Expected scented (0,0) to be true but got %+v", sm.planetScent[0][0])
	}
}
