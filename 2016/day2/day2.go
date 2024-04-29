package main

import (
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {

	lines, err := readTextFile("day2.txt")

	if err != nil {
		fmt.Println(err)
	}

	w := walker{
		vector{1, 1},
		[][]rune{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		2,
	}

	bathroomCode := []rune{}

	for _, line := range lines {
		for _, direction := range line {
			w.move(direction)
		}
		bathroomCode = append(bathroomCode, w.keypad[w.currentLocation.y][w.currentLocation.x])
	}

	fmt.Println(w)
	fmt.Println(bathroomCode)
}
func part2() {

	lines, err := readTextFile("day2.txt")

	if err != nil {
		fmt.Println(err)
	}

	w := walker{
		vector{1, 1},
		[][]rune{
			{'#', '#', '1', '#', '#'},
			{'#', '2', '3', '4', '#'},
			{'5', '6', '7', '8', '9'},
			{'#', 'A', 'B', 'C', '#'},
			{'#', '#', 'D', '#', '#'},
		},
		4,
	}

	bathroomCode := []string{}

	for _, line := range lines {
		for _, direction := range line {
			w.move(direction)
		}
		bathroomCode = append(bathroomCode, string(w.keypad[w.currentLocation.y][w.currentLocation.x]))
	}

	fmt.Println(w)
	fmt.Println(bathroomCode)
}
