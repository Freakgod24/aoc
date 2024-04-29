package main

import (
	"fmt"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {

	line, err := readTextFile("day1.txt")

	if err != nil {
		fmt.Println(err)
	}

	w := walker{
		vector{0, 1},
		vector{0, 0},
		make(map[vector]bool),
		vector{0, 0},
	}

	instructions := strings.Split(strings.ReplaceAll(line[0], " ", ""), ",")

	for _, instruction := range instructions {
		w.parseInstruction(instruction, false)
	}

	fmt.Println(abs(w.currentLocation.x) + abs(w.currentLocation.y))
}
func part2() {

	line, err := readTextFile("day1.txt")

	if err != nil {
		fmt.Println(err)
	}

	w := walker{
		vector{0, 1},
		vector{0, 0},
		make(map[vector]bool),
		vector{0, 0},
	}

	instructions := strings.Split(strings.ReplaceAll(line[0], " ", ""), ",")

	for _, instruction := range instructions {
		if w.parseInstruction(instruction, true) {
			fmt.Println(abs(w.currentLocation.x) + abs(w.currentLocation.y))
			break
		}
	}
}
