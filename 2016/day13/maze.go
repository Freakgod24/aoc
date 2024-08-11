package main

import (
	"fmt"
	"math/bits"
)

type Location struct {
	x       uint
	y       uint
	g       uint
	h       uint
	f       uint
	wall    bool
	visited bool
	path    bool
	parent  *Location
}

type Maze struct {
	width          uint
	height         uint
	favoriteNumber uint
	grid           [][]Location
}

func (maze *Maze) init(width uint, height uint, favoriteNumber uint) {
	maze.favoriteNumber = favoriteNumber
	maze.width = width
	maze.height = height
	maze.grid = make([][]Location, height)

	for y := range maze.height {
		maze.grid[y] = make([]Location, width)
		for x := range maze.width {
			maze.grid[y][x] = Location{x: x, y: y, wall: maze.isWall(x, y)}
		}
	}
}

func (maze *Maze) clear() {
	for y := range maze.height {
		for x := range maze.width {
			maze.grid[y][x].f = 0
			maze.grid[y][x].g = 0
			maze.grid[y][x].h = 0
			maze.grid[y][x].visited = false
			maze.grid[y][x].path = false
			maze.grid[y][x].parent = nil
		}
	}
}

func (maze *Maze) isWall(x uint, y uint) bool {
	sum := x*x + 3*x + 2*x*y + y + y*y + maze.favoriteNumber
	return bits.OnesCount(sum)%2 == 1
}

func (maze *Maze) getAvailableDirections(currentLocation *Location) []*Location {
	var locations []*Location

	directions := [4][2]int{
		{0, 1},  //Up
		{0, -1}, //Down
		{-1, 0}, //Left
		{1, 0},  //Right
	}

	for _, direction := range directions {
		newX := int(currentLocation.x) + direction[0]
		newY := int(currentLocation.y) + direction[1]

		if newX >= 0 && newX < int(maze.width) && newY >= 0 && newY < int(maze.height) {
			if maze.grid[newY][newX].wall == false {
				locations = append(locations, &maze.grid[newY][newX])
			}
		}
	}

	return locations
}

func (maze *Maze) print() {
	fmt.Println()
	steps := -1
	for y := range maze.height {
		for x := range maze.width {
			cell := maze.grid[y][x]
			if cell.wall {
				fmt.Print("#")
			} else if cell.path {
				fmt.Print("O")
				steps++
			} else if cell.visited {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	fmt.Println()
}

func (maze *Maze) solve(startLocation *Location, endLocation *Location) (bool, int) {
	var searchLocations = PriorityQueue{}
	searchLocations.push(startLocation, 0)

	for searchLocations.root != nil {
		var currentLocation = searchLocations.pop().(*Location)

		if currentLocation == endLocation {
			// A solution was found
			steps := 0
			for currentLocation.parent != nil {
				currentLocation.path = true
				currentLocation = currentLocation.parent
				steps++
			}
			currentLocation.path = true
			return true, steps
		}

		currentLocation.visited = true

		for _, nextLocation := range maze.getAvailableDirections(currentLocation) {

			if nextLocation.visited {
				continue
			}

			newG := currentLocation.g + 1
			newH := getHeuristicDistance(nextLocation, endLocation)
			newF := newG + newH

			inSearchLocations := searchLocations.contains(nextLocation)

			if !inSearchLocations || (inSearchLocations && newG < nextLocation.g) {
				nextLocation.g = newG
				nextLocation.f = newF
				nextLocation.parent = currentLocation
			}

			if !inSearchLocations {
				searchLocations.push(nextLocation, nextLocation.f)
			}
		}
	}

	return false, -1
}

func getHeuristicDistance(fromLocation *Location, toLocation *Location) uint {
	dx := toLocation.x - fromLocation.x
	dy := toLocation.y - fromLocation.y

	// Return the sum of dx and dy. It is not the exact euclidean distance,
	// but more of a proportial metric. No real added value to compute
	// square root and get the exact value.
	return dx + dy
}
