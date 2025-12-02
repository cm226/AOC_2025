package main

import (
	"fmt"
	"strconv"

	"AOC_2025/util"
)

func main() {
	commands := []string{}
	util.Read(func(line string) {
		commands = append(commands, line)
	}, "day1/input.txt")

	fmt.Println("Part1: ", part1(commands))
	fmt.Println("part2: ", part2(commands)+part1(commands))
}

func part2(commands []string) int {
	position := 50
	timesAt0 := 0
	for _, command := range commands {
		direction := 1
		if command[0] != 'R' {
			direction = -1
		}

		ammount, err := strconv.Atoi(command[1:])
		if err != nil {
			panic(err)
		}
		position += ammount * direction
		for position < 0 {
			position = position + 100
			timesAt0 += 1
		}
		for position > 99 {
			position = position - 100
			timesAt0 += 1
		}

		if position == 0 {
			timesAt0 -= 1
		}
	}
	return timesAt0
}

func part1(commands []string) int {
	position := 50
	timesAt0 := 0
	for _, command := range commands {
		direction := 1
		if command[0] != 'R' {
			direction = -1
		}

		ammount, err := strconv.Atoi(command[1:])
		if err != nil {
			panic(err)
		}
		position += ammount * direction
		for position < 0 {
			position = position + 100
		}
		for position > 99 {
			position = position - 100
		}

		if position == 0 {
			timesAt0 += 1
		}
	}
	return timesAt0
}
