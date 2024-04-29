package main

import (
	"fmt"
	"strconv"
)

type walker struct {
	facingDirection  vector
	currentLocation  vector
	visitedLocations map[vector]bool
	visitedTwice     vector
}

func (w *walker) turnLeft() {
	w.facingDirection.rotateMinus90()
}

func (w *walker) turnRight() {
	w.facingDirection.rotatePlus90()
}

func (w *walker) moveForward(steps int, stopAtVisitedTwice bool) bool {
	for s := 1; s <= steps; s++ {
		w.currentLocation.x += w.facingDirection.x
		w.currentLocation.y += w.facingDirection.y

		if stopAtVisitedTwice && w.visitedLocations[w.currentLocation] == true {
			w.visitedTwice = w.currentLocation
			return true
		} else {
			w.visitedLocations[w.currentLocation] = true
		}
	}

	return false
}

func (w *walker) parseInstruction(instruction string, stopAtVisitedTwice bool) bool {

	turn := instruction[0]
	steps, err := strconv.Atoi(string(instruction[1:]))
	if err != nil {
		panic(fmt.Sprintf("invalid steps received: %s", instruction[1:]))
	}

	switch turn {
	case 'L':
		w.turnLeft()
	case 'R':
		w.turnRight()
	default:
		panic(fmt.Sprintf("invalid turn received: %c", turn))
	}

	return w.moveForward(steps, stopAtVisitedTwice)
}
