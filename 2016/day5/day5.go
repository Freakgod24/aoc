package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	test()
	part1()
	part2()
}

func test() {
	max_length := 8
	input_str := "abc"

	index := 0
	hash := ""
	password := ""
	current_length := 0

	for current_length < max_length {
		door_id := []byte(input_str + strconv.Itoa(index))
		hash = fmt.Sprintf("%x", md5.Sum(door_id))

		if strings.HasPrefix(hash, "00000") {
			password += string(hash[5])
			current_length += 1
		}

		index += 1
	}

	fmt.Println(password)
}

func part1() {
	max_length := 8
	input_str := "cxdnnyjw"

	index := 0
	hash := ""
	password := ""
	current_length := 0

	for current_length < max_length {
		door_id := []byte(input_str + strconv.Itoa(index))
		hash = fmt.Sprintf("%x", md5.Sum(door_id))

		if strings.HasPrefix(hash, "00000") {
			password += string(hash[5])
			current_length += 1
		}

		index += 1
	}

	fmt.Println(password)
}

func part2() {
	max_length := 8
	input_str := "cxdnnyjw"

	index := 0
	hash := ""
	password := [8]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	current_length := 0

	for current_length < max_length {
		door_id := []byte(input_str + strconv.Itoa(index))
		hash = fmt.Sprintf("%x", md5.Sum(door_id))
		current_length = 0

		if strings.HasPrefix(hash, "00000") {
			if '0' <= hash[5] && hash[5] <= '7' {
				position, _ := strconv.Atoi(string(hash[5]))
				char := hash[6]

				// Do not overwrite existing characters
				if password[position] == ' ' {
					password[position] = char
					fmt.Printf("%s\n", password)
				}
			}

			// Count non-space characters. Need 8 to decrypt the password.
			for _, c := range password {
				if c != ' ' {
					current_length += 1
				}
			}
		}

		index += 1
	}
}
