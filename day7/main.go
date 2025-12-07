package main

import (
	"fmt"

	"AOC_2025/util"
)

func split(grid [][]rune, x int, y int) bool {
	return grid[y][x] == '^'
}

func runSimulationPt1(grid [][]rune, tachions map[int]int, y int) (map[int]int, int) {
	splits := 0
	for x, v := range tachions {
		if v != 0 && split(grid, x, y) {
			tachions[x-1] += tachions[x]
			tachions[x+1] += tachions[x]
			tachions[x] = 0
			splits += 1
		}
	}
	return tachions, splits
}

func main() {
	grid := util.ReadMatrix("day7/input.txt")
	tachions := map[int]int{}
	for i, c := range grid[0] {
		if c == 'S' {
			tachions[i] = 1
			break
		}
	}

	totalSplits := 0
	for y := range len(grid) - 1 {
		splits := 0
		tachions, splits = runSimulationPt1(grid, tachions, y)
		totalSplits += splits
	}

	fmt.Println("pt1", totalSplits)

	timelines := 0
	for _, v := range tachions {
		timelines += v
	}
	fmt.Println("pt2: ", timelines)
}
