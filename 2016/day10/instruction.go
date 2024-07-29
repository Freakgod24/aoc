package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ----------------------
// Type definitions
// ----------------------
type Location int

const (
	LocationNone   Location = -1
	LocationInput  Location = 0
	LocationBot    Location = 1
	LocationOutput Location = 2
)

type Index int

const (
	IndexNone Index = -1
)

type InstructionType int

const (
	InstructionGoes  InstructionType = 0
	InstructionGives InstructionType = 1
)

type Instruction struct {
	instructionType     InstructionType
	fromLocation        Location
	fromLocationIndex   Index
	toLocationLow       Location
	toLocationLowIndex  Index
	toLocationHigh      Location
	toLocationHighIndex Index
}

// ----------------------
// Errors definitions
// ----------------------
var ErrInvalidInstruction = fmt.Errorf("invalid instruction ")
var ErrInvalidNumberOfArguments = fmt.Errorf("invalid number of arguments ")
var ErrInvalidDestination = fmt.Errorf("invalid destination ")
var ErrInvalidIndex = fmt.Errorf("invalid index ")

// ----------------------
// Type implementations
// ----------------------
func parseInstruction(instructionStr string) (Instruction, error) {
	args := strings.Split(instructionStr, " ")

	switch args[0] {
	case "value":

		if len(args) != 6 {
			return Instruction{}, ErrInvalidNumberOfArguments
		} else {
			return parseGoesCommand(args)
		}

	case "bot":
		if len(args) != 12 {
			return Instruction{}, ErrInvalidNumberOfArguments
		} else {
			return parseGivesCommand(args)
		}

	default:
		return Instruction{}, ErrInvalidInstruction
	}
}

func parseGoesCommand(args []string) (Instruction, error) {

	locationIndex, err1 := parseIndex(args[1])
	locationLow, err2 := parseLocation(args[4])
	locationLowIndex, err3 := parseIndex(args[5])

	if err1 != nil {
		return Instruction{}, err1
	} else if err2 != nil {
		return Instruction{}, err2
	} else if err3 != nil {
		return Instruction{}, err3
	} else {
		return Instruction{
			InstructionGoes,
			LocationInput,
			locationIndex,
			locationLow,
			locationLowIndex,
			LocationNone,
			IndexNone,
		}, nil
	}
}

func parseGivesCommand(args []string) (Instruction, error) {

	locationIndex, err1 := parseIndex(args[1])
	locationLow, err2 := parseLocation(args[5])
	locationLowIndex, err3 := parseIndex(args[6])
	locationHigh, err4 := parseLocation(args[10])
	locationHighIndex, err5 := parseIndex(args[11])

	if err1 != nil {
		return Instruction{}, err1
	} else if err2 != nil {
		return Instruction{}, err2
	} else if err3 != nil {
		return Instruction{}, err3
	} else if err4 != nil {
		return Instruction{}, err4
	} else if err5 != nil {
		return Instruction{}, err5
	} else {
		return Instruction{
			InstructionGives,
			LocationBot,
			locationIndex,
			locationLow,
			locationLowIndex,
			locationHigh,
			locationHighIndex,
		}, nil
	}
}

func parseLocation(destinationStr string) (Location, error) {
	switch destinationStr {
	case "bot":
		return LocationBot, nil
	case "output":
		return LocationOutput, nil
	default:
		return LocationNone, ErrInvalidDestination
	}
}

func parseIndex(indexStr string) (Index, error) {
	indexValue, err := strconv.Atoi(indexStr)
	if err == nil {
		return Index(indexValue), nil
	} else {
		return IndexNone, ErrInvalidIndex
	}
}
