package main

import (
	"fmt"
	"math"
)

func main() {
	test()
	part1()
	part2()
}

func test() {
	messages := []string{
		"eedadn",
		"drvtee",
		"eandsr",
		"raavrd",
		"atevrs",
		"tsrnev",
		"sdttsa",
		"rasrtv",
		"nssdts",
		"ntnada",
		"svetve",
		"tesnvt",
		"vntsnd",
		"vrdear",
		"dvrsen",
		"enarar",
	}

	message_corrected_max := ""
	message_corrected_min := ""

	// Initialize a data structure to store the characters repetitions
	// The structure will be composed of six map, one for each position.
	var char_repetitions [6]map[rune]int
	for i := range char_repetitions {
		char_repetitions[i] = make(map[rune]int)
	}

	// Parse each message and index all the character repetition for
	// each position
	for _, message := range messages {
		for i, c := range message {
			char_repetitions[i][c] += 1
		}
	}

	// Extract the most repeated character for each position
	for _, values := range char_repetitions {
		message_corrected_max += get_max_repeated_char(values)
		message_corrected_min += get_min_repeated_char(values)
	}

	fmt.Println("Corrected message (max):", message_corrected_max)
	fmt.Println("Corrected message (min):", message_corrected_min)
}

func get_max_repeated_char(values map[rune]int) string {
	max_rune := ' '
	max_int := 0

	for r, i := range values {
		if i > max_int {
			max_int = i
			max_rune = r
		}
	}

	return string(max_rune)
}

func get_min_repeated_char(values map[rune]int) string {
	min_rune := ' '
	min_int := math.MaxInt32

	for r, i := range values {
		if i < min_int {
			min_int = i
			min_rune = r
		}
	}

	return string(min_rune)
}

func part1() {
	// Here, we are assuming ANSI encoded text file. 1 Byte per character.
	messages, _ := readTextFile("day6.txt")
	message_corrected := ""

	var char_repetitions [8]map[rune]int
	for i := range char_repetitions {
		char_repetitions[i] = make(map[rune]int)
	}

	for _, message := range messages {
		for i, c := range message {
			char_repetitions[i][c] += 1
		}
	}

	for _, values := range char_repetitions {
		message_corrected += get_max_repeated_char(values)
	}

	fmt.Println("Corrected message:", message_corrected)
}

func part2() {
	// Here, we are assuming ANSI encoded text file. 1 Byte per character.
	messages, _ := readTextFile("day6.txt")
	message_corrected := ""

	var char_repetitions [8]map[rune]int
	for i := range char_repetitions {
		char_repetitions[i] = make(map[rune]int)
	}

	for _, message := range messages {
		for i, c := range message {
			char_repetitions[i][c] += 1
		}
	}

	for _, values := range char_repetitions {
		message_corrected += get_min_repeated_char(values)
	}

	fmt.Println("Corrected message:", message_corrected)
}
