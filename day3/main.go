package main

import (
	"fmt"

	"AOC_2025/util"
)

func findMaxin(subBank string) int {
	maxV := 0
	maxI := 0
	for i := range len(subBank) {
		v1 := int(subBank[i] - '0')
		if v1 > maxV {
			maxV = v1
			maxI = i
		}
	}
	return maxI
}

func pt2(bank string, n int) int {
	maxNumIndexes := []int{}
	prevI := -1
	for i := range n {
		prevI = findMaxin(bank[prevI+1:len(bank)-(n-i-1)]) + prevI + 1
		maxNumIndexes = append(maxNumIndexes, prevI)
	}

	max := 0
	pow := 1
	for i := len(maxNumIndexes) - 1; i >= 0; i-- {
		v := int(bank[maxNumIndexes[i]]-'0') * pow
		pow *= 10
		max += v
	}
	return max
}

func main() {
	part1Sum := 0
	part2Sum := 0
	util.Read(func(line string) {
		fmt.Println("line", line)
		part1Sum += pt2(line, 2)
		part2Sum += pt2(line, 12)
	}, "./day3/input.txt")

	fmt.Println("Part1: ", part1Sum)
	fmt.Println("Part2: ", part2Sum)
}
