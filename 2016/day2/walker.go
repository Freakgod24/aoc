package main

import (
	"fmt"
)

type walker struct {
	currentLocation vector
	keypad          [][]rune
	keypadSize      int
}

func (w *walker) moveLeft() {
	if w.currentLocation.x > 0 && w.keypad[w.currentLocation.y][w.currentLocation.x-1] != '#' {
		w.currentLocation.x--
	}
}

func (w *walker) moveRight() {
	if w.currentLocation.x < w.keypadSize && w.keypad[w.currentLocation.y][w.currentLocation.x+1] != '#' {
		w.currentLocation.x++
	}
}

func (w *walker) moveUp() {
	if w.currentLocation.y > 0 && w.keypad[w.currentLocation.y-1][w.currentLocation.x] != '#' {
		w.currentLocation.y--
	}
}

func (w *walker) moveDown() {
	if w.currentLocation.y < w.keypadSize && w.keypad[w.currentLocation.y+1][w.currentLocation.x] != '#' {
		w.currentLocation.y++
	}
}

func (w *walker) move(direction rune) {
	switch direction {
	case 'L':
		w.moveLeft()
	case 'R':
		w.moveRight()
	case 'U':
		w.moveUp()
	case 'D':
		w.moveDown()
	default:
		panic(fmt.Sprintf("invalid direction received: %c", direction))
	}

	fmt.Printf("%c %d\n", direction, w.keypad[w.currentLocation.y][w.currentLocation.x])
}
