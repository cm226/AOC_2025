package main

import (
	"fmt"

	"AOC_2025/util"

	"github.com/s0rg/vec2d"
)

func tr(m [][]rune, pos vec2d.V[int]) bool {
	if pos.Y < 0 || pos.X < 0 || pos.X >= len(m[0]) || pos.Y >= len(m) {
		return false
	}
	return m[pos.Y][pos.X] == '@'
}

func runSimulation(grid [][]rune, moveRoles bool) int {
	up := vec2d.New(0, -1)
	down := vec2d.New(0, 1)
	left := vec2d.New(-1, 0)
	right := vec2d.New(1, 0)

	all := []vec2d.V[int]{
		up, down, left, right,
		up.Add(left), up.Add(right), down.Add(left), down.Add(right),
	}

	moveable := 0
	for y := range grid {
		for x := range grid[y] {
			current := vec2d.New(x, y)

			if tr(grid, current) {
				trCount := 0
				for _, pos := range all {
					if tr(grid, current.Add(pos)) {
						trCount += 1
					}
				}
				if trCount < 4 {
					moveable += 1
					if moveRoles {
						grid[current.Y][current.X] = '.'
					}
				}
			}
		}
	}
	return moveable
}

func main() {
	grid := util.ReadMatrix("day4/input.txt")
	pt1 := runSimulation(grid, false)
	fmt.Println("Part1: ", pt1)

	pt2 := -1
	pt2Total := 0
	for pt2 != 0 {
		pt2 = runSimulation(grid, true)
		pt2Total += pt2
	}
	fmt.Println("Part 2: ", pt2Total)
}
