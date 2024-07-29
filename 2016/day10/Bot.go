package main

type Bot struct {
	chips       []Chip
	instruction Instruction
}

func (bot *Bot) getLowChip() Chip {
	if bot.chips[0].value < bot.chips[1].value {
		return bot.chips[0]
	} else {
		return bot.chips[1]
	}
}

func (bot *Bot) getHighChip() Chip {
	if bot.chips[0].value > bot.chips[1].value {
		return bot.chips[0]
	} else {
		return bot.chips[1]
	}
}
