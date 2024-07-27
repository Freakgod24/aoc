package main

import (
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRectSize = fmt.Errorf("Invalid size for rect cmd")
var ErrInvalidDirection = fmt.Errorf("Invalid direction for rotate cmd")
var ErrInvalidPosition = fmt.Errorf("Invalid position for rotate cmd")
var ErrInvalidAmount = fmt.Errorf("Invalid amount for rotate cmd")

const SCREEN_WIDTH = 50
const SCREEN_HEIGHT = 6

type Pixels [SCREEN_HEIGHT][SCREEN_WIDTH]string
type Screen struct {
	pixels Pixels
}

func (s *Screen) width() int {
	return SCREEN_WIDTH
}

func (s *Screen) height() int {
	return SCREEN_HEIGHT
}

func main() {
	// test()
	part1()
	// part2()
}

// func test() {
// 	var err error
// 	screen := Screen{Pixels{
// 		{".", ".", ".", ".", ".", ".", "."},
// 		{".", ".", ".", ".", ".", ".", "."},
// 		{".", ".", ".", ".", ".", ".", "."}}}
//
// 	instructions := []string{
// 		"rect 3x2",
// 		"rotate column x=1 by 1",
// 		"rotate row y=0 by 4",
// 		"rotate column x=1 by 1",
// 	}
//
// 	for _, ins := range instructions {
// 		tokens := strings.Split(ins, " ")
// 		cmd := tokens[0]
//
// 		switch cmd {
// 		case "rect":
// 			_, err = executeRect(&screen, tokens[1])
// 		case "rotate":
// 			_, err = executeRotate(&screen, tokens[1], tokens[2], tokens[4])
// 		}
//
// 		if err != nil {
// 			fmt.Println(err)
// 		}
//
// 		fmt.Println("")
// 		for _, row := range screen.pixels {
// 			fmt.Println(row)
// 		}
//
// 	}
//
// }

func part1() {
	var err error
	var pixels_lit int

	// Initialize the screen with the default pixels
	screen := Screen{}
	for row := range screen.height() {
		for col := range screen.width() {
			screen.pixels[row][col] = "."
		}
	}

	instructions, _ := readTextFile("day8.txt")

	for _, ins := range instructions {
		tokens := strings.Split(ins, " ")
		cmd := tokens[0]

		switch cmd {
		case "rect":
			_, err = executeRect(&screen, tokens[1])
		case "rotate":
			_, err = executeRotate(&screen, tokens[1], tokens[2], tokens[4])
		}

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("")
		for _, row := range screen.pixels {
			fmt.Println(row)
		}

	}
	
	for _, row := range screen.pixels {
		for _, pixel := range row {
			if pixel == "#" {
				pixels_lit += 1
			}
		}
	}

	fmt.Println(pixels_lit)
}

func executeRect(screen *Screen, sizeStr string) (*Screen, error) {
	// Split the size string, expecting format WxH
	dimensions := strings.Split(sizeStr, "x")
	if len(dimensions) != 2 {
		return screen, ErrInvalidRectSize
	}

	// Parse and validate the width and height
	width, err1 := strconv.Atoi(dimensions[0])
	height, err2 := strconv.Atoi(dimensions[1])
	if err1 != nil || err2 != nil || width > screen.width() || height > screen.height() || width <= 0 || height <= 0 {
		return screen, ErrInvalidRectSize
	}

	// Draw the rectangle on the screen
	for y := range height {
		for x := range width {
			screen.pixels[y][x] = "#"
		}
	}

	return screen, nil
}

func executeRotate(screen *Screen, dirStr string, posStr string, amountStr string) (*Screen, error) {
	// Parse the position string, expecting format x=1 or y=1
	posArr := strings.Split(posStr, "=")
	if len(posArr) != 2 {
		return screen, ErrInvalidPosition
	}

	pos, err := strconv.Atoi(posArr[1])
	if err != nil || pos < 0 {
		return screen, ErrInvalidPosition
	}

	// Parse the amount string, expecting a single number.
	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount <= 0 {
		return screen, ErrInvalidAmount
	}

	// Determine the direction and apply rotation
	switch dirStr {
	case "column":

		if pos >= screen.width() {
			return screen, ErrInvalidPosition
		}

		executeRotateColumn(screen, pos, amount)

	case "row":

		if pos >= screen.height() {
			return screen, ErrInvalidPosition
		}

		executeRotateRow(screen, pos, amount)

	default:
		return screen, ErrInvalidDirection
	}

	return screen, nil
}

func executeRotateColumn(screen *Screen, col int, amount int) (*Screen, error) {
	newColumn := make([]string, screen.height())

	// Calculate the new positions and pixels
	for row := range screen.height() {
		newPosition := (row + amount) % screen.height()
		newColumn[newPosition] = screen.pixels[row][col]
	}

	// Draw the updated column on the screen
	for row, pixel := range newColumn {
		screen.pixels[row][col] = pixel
	}

	return screen, nil
}

func executeRotateRow(screen *Screen, row int, amount int) (*Screen, error) {
	newRow := make([]string, screen.width())

	// Calculate the new positions and pixels
	for col := range screen.width() {
		newPosition := (col + amount) % screen.width()
		newRow[newPosition] = screen.pixels[row][col]
	}

	// Draw the updated row on the screen
	for col, pixel := range newRow {
		screen.pixels[row][col] = pixel
	}

	return screen, nil
}
