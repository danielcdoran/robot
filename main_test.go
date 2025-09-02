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

	expected := NewRobotDetails(pos, South)
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}

func TestTurnRightStartNorth(t *testing.T) {
	pos := NewRobotPosition(10, 10)
	northFacing := NewRobotDetails(pos, North)
	result := turnRight(northFacing)

	expected := NewRobotDetails(pos, East)
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}

func TestTurnRightStartEast(t *testing.T) {
	pos := NewRobotPosition(10, 10)
	eastFacing := NewRobotDetails(pos, East)
	result := turnRight(eastFacing)

	expected := NewRobotDetails(pos, South)
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}

func TestTurnRightStartSouth(t *testing.T) {
	pos := NewRobotPosition(10, 10)
	southFacing := NewRobotDetails(pos, South)
	result := turnRight(southFacing)

	expected := NewRobotDetails(pos, West)
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}

func TestTurnRightStartWest(t *testing.T) {
	pos := NewRobotPosition(10, 10)
	westFacing := NewRobotDetails(pos, West)
	result := turnRight(westFacing)

	expected := NewRobotDetails(pos, North)
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}

func TestMoveNorth(t *testing.T) {
	pos := NewRobotPosition(10, 10)
	northFacing := NewRobotDetails(pos, North)
	result := moveInDirection(northFacing)

	expectedPos := NewRobotPosition(10, 11)
	expected := NewRobotDetails(expectedPos, North)
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}

func TestMoveEast(t *testing.T) {
	pos := NewRobotPosition(10, 10)
	eastFacing := NewRobotDetails(pos, East)
	result := moveInDirection(eastFacing)
	expectedPos := NewRobotPosition(11, 10)
	expected := NewRobotDetails(expectedPos, East)
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}

func TestMoveSouth(t *testing.T) {
	pos := NewRobotPosition(10, 10)
	southFacing := NewRobotDetails(pos, South)
	result := moveInDirection(southFacing)
	expectedPos := NewRobotPosition(10, 9)
	expected := NewRobotDetails(expectedPos, South)
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}

func TestMoveWest(t *testing.T) {
	pos := NewRobotPosition(10, 10)
	westFacing := NewRobotDetails(pos, West)
	result := moveInDirection(westFacing)
	expectedPos := NewRobotPosition(9, 10)
	expected := NewRobotDetails(expectedPos, West)
	if expected != result {
		t.Errorf("Expected Left turn result %+v different from result  position %+v", expected, result)
	}
}
