package main

import (
	"fmt"
	// "strconv"
	// "strings"
)

var bots [255]Bot
var outputs [255]Chip

func main() {
	part1()
	part2()
}

func part1() {
	lines, _ := readTextFile("day10.txt")
	var nextBots []*Bot

	// First pass. Gives chip and instruction to corresponding bots.
	for _, line := range lines {
		instruction, err := parseInstruction(line)

		if err != nil {
			fmt.Println(err)
			return
		}

		switch instruction.instructionType {
		case InstructionGoes:
			bot := &bots[instruction.toLocationLowIndex]
			chips := &bot.chips
			value := int(instruction.fromLocationIndex)
			*chips = append(*chips, Chip{value})

			// Store a list of all bots that were given 2 chips already
			// Those will be the starting point of the simulation.
			if len(*chips) == 2 {
				nextBots = append(nextBots, bot)
				fmt.Println("Bot", instruction.toLocationLowIndex, "have 2 chips")
			}

		case InstructionGives:
			bots[instruction.fromLocationIndex].instruction = instruction
		}
	}

	// Perform simulation until no bots are having two chips
	for len(nextBots[0].chips) == 2 {
		bot := nextBots[0]

		lowChip := bot.getLowChip()
		highChip := bot.getHighChip()

		lowIndex := bot.instruction.toLocationLowIndex
		highIndex := bot.instruction.toLocationHighIndex

		// Look if the end condition is met
		// otherwise proceed with the simulation
		if lowChip.value == 17 && highChip.value == 61 {
			fmt.Println("Condition found")
			fmt.Println(bot.instruction.fromLocationIndex)
			// break
		}

		// Gives the lowest-value chip
		switch bot.instruction.toLocationLow {
		case LocationBot:
			bots[lowIndex].chips = append(bots[lowIndex].chips, lowChip)
			fmt.Println("Giving chip", lowChip.value, "to bot", lowIndex)
		case LocationOutput:
			outputs[lowIndex].value = lowChip.value
			fmt.Println("Giving chip", lowChip.value, "to output", lowIndex)
		}

		// Gives the highest-value chip
		switch bot.instruction.toLocationHigh {
		case LocationBot:
			bots[highIndex].chips = append(bots[highIndex].chips, highChip)
			fmt.Println("Giving chip", highChip.value, "to bot", highIndex)
		case LocationOutput:
			outputs[highIndex].value = highChip.value
			fmt.Println("Giving chip", highChip.value, "to bot", highIndex)
		}

		// After giving all the chips to other bots
		// Remove this bot references to those chips
		// and remove this bot from nextBots
		bot.chips = nil
		nextBots = nextBots[1:]

		// Look if any of those two bots are now having
		// two chips. In theory, at least one of those bots
		// should have two chips. If not, the simulation will end
		if len(bots[lowIndex].chips) == 2 {
			nextBots = append(nextBots, &bots[lowIndex])
		}
		if len(bots[highIndex].chips) == 2 {
			nextBots = append(nextBots, &bots[highIndex])
		}
	}
}

func part2() {
	value := outputs[0].value * outputs[1].value * outputs[2].value
	fmt.Println(value)
}
