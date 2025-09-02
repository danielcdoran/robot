package main

import (
	"testing"
)

func TestTurnLeftStartNorth(t *testing.T) {
	pos := NewRobotPosition(10, 10)
	northFacing := NewRobotDetails(pos, North)
	result := turnLeft(northFacing)

	expected := NewRobotDetails(pos, West)
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}

func TestTurnLeftStartEast(t *testing.T) {
	pos := NewRobotPosition(10, 10)
	eastFacing := NewRobotDetails(pos, East)
	result := turnLeft(eastFacing)

	expected := NewRobotDetails(pos, North)
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}

func TestTurnLeftStartSouth(t *testing.T) {
	pos := NewRobotPosition(10, 10)
	southFacing := NewRobotDetails(pos, South)
	result := turnLeft(southFacing)

	expected := NewRobotDetails(pos, East)
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}

func TestTurnLeftStartWest(t *testing.T) {
	pos := NewRobotPosition(10, 10)
	westFacing := NewRobotDetails(pos, West)
	result := turnLeft(westFacing)

	expected := NewRobotDetails(pos, West)
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}
