package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Computer struct {
	a        int
	b        int
	c        int
	d        int
	program  []string
	location int
}

func (computer *Computer) cpy(source string, destination string) error {
	var value int

	switch source {
	case "a":
		value = computer.a
	case "b":
		value = computer.b
	case "c":
		value = computer.c
	case "d":
		value = computer.d
	default:
		intLiteral, err := strconv.Atoi(source)
		if err != nil {
			return fmt.Errorf("Invalid source register or value")
		} else {
			value = intLiteral
		}
	}

	switch destination {
	case "a":
		computer.a = value
	case "b":
		computer.b = value
	case "c":
		computer.c = value
	case "d":
		computer.d = value
	default:
		return fmt.Errorf("Invalid destination register")
	}

	return nil
}

func (computer *Computer) inc(register string) error {
	switch register {
	case "a":
		computer.a++
	case "b":
		computer.b++
	case "c":
		computer.c++
	case "d":
		computer.d++
	default:
		return fmt.Errorf("Invalid register")
	}

	return nil
}

func (computer *Computer) dec(register string) error {
	switch register {
	case "a":
		computer.a--
	case "b":
		computer.b--
	case "c":
		computer.c--
	case "d":
		computer.d--
	default:
		return fmt.Errorf("Invalid Register")
	}

	return nil
}

func (computer *Computer) jnz(condition string, offset string) error {
	value, _ := strconv.Atoi(offset)
	jump := false

	switch condition {
	case "a":
		jump = computer.a != 0
	case "b":
		jump = computer.b != 0
	case "c":
		jump = computer.c != 0
	case "d":
		jump = computer.d != 0
	default:
		intLiteral, err := strconv.Atoi(condition)
		if err == nil {
			jump = intLiteral != 0
		} else {
			return fmt.Errorf("Invalid jump condition")
		}
	}

	if jump {
		computer.location += value - 1
	}

	return nil
}

func (computer *Computer) run() error {
	var err error

	for computer.location < len(computer.program) {
		programLine := computer.program[computer.location]
		tokens := strings.Split(programLine, " ")

		switch tokens[0] {
		case "cpy":
			err = computer.cpy(tokens[1], tokens[2])
		case "inc":
			err = computer.inc(tokens[1])
		case "dec":
			err = computer.dec(tokens[1])
		case "jnz":
			err = computer.jnz(tokens[1], tokens[2])
		}

		if err != nil {
			return fmt.Errorf("Program halted")
		}

		computer.location++
	}

	return nil
}

func main() {
	test()
	part1()
	part2()
}

func test() {
	computer := Computer{program: []string{
		"cpy 41 a",
		"inc a",
		"inc a",
		"dec a",
		"jnz a 2",
		"dec a",
	}}
	computer.run()
	fmt.Println(computer)
}

func part1() {
	programLines, _ := readTextFile("day12.txt")
	computer := Computer{program: programLines}
	computer.run()
	fmt.Println(computer)
}

func part2() {
	programLines, _ := readTextFile("day12.txt")
	computer := Computer{c: 1, program: programLines}
	computer.run()
	fmt.Println(computer)
}
