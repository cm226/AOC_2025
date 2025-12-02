package main

import (
	"fmt"
	"strconv"
	"strings"

	"AOC_2025/util"
)

type IDRange struct {
	start int
	end   int
}

func isRepeatPart1(number int) bool {
	str := strconv.Itoa(number)
	if len(str)%2 != 0 {
		return false
	}

	return str[len(str)/2:] == str[:len(str)/2]
}

func isRepeatPart2(number int) bool {
	str := strconv.Itoa(number)

	splits := 2
out:
	for splits <= len(str) {
		if len(str)%splits != 0 {
			splits += 1
			continue
		}
		for i := range splits {
			splitSize := len(str) / splits
			if str[0:splitSize] != str[splitSize*i:splitSize*i+splitSize] {
				splits += 1
				continue out
			}
		}
		return splits <= len(str)
	}

	return false
}

func main() {
	isRepeatPart2(1000)
	idRanges := []IDRange{}

	util.ReadCSV(func(line string) {
		line = strings.TrimSpace(line)
		startEnd := strings.Split(line, "-")

		start, err := strconv.Atoi(startEnd[0])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(startEnd[1])
		if err != nil {
			panic(err)
		}

		idRanges = append(idRanges, IDRange{start, end})
	},
		"day2/input.txt")

	sum := 0
	for _, rng := range idRanges {
		current := rng.start
		for current <= rng.end {
			if isRepeatPart2(current) {
				fmt.Println(current)
				sum += current
			}
			current += 1
		}
	}
	fmt.Println("Sum : ", sum)
}
