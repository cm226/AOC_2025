package main

import (
	"fmt"
	"strconv"
	"strings"

	"AOC_2025/util"
)

func doTheMath(collumns [][]int, ops []string) int {
	answers := []int{}
	for i := range collumns {
		answers = append(answers, collumns[i][0])
	}

	for j, column := range collumns {
		for _, stVal := range column[1:] {
			if ops[j] == "+" {
				answers[j] = answers[j] + stVal
			}

			if ops[j] == "*" {
				answers[j] = answers[j] * stVal
			}
		}
	}

	sum := 0
	for _, a := range answers {
		sum += a
	}
	return sum
}

func toIntColumns(cols [][]string) [][]int {
	intCols := [][]int{}
	for i := range cols {
		if cols[i] == nil {
			continue
		}
		intCols = append(intCols, []int{})
		for j := range cols[i] {
			trimmed := strings.TrimSpace(cols[i][j])
			if trimmed == "" {
				continue
			}
			v, err := strconv.Atoi(trimmed)
			if err != nil {
				panic(err)
			}
			intCols[i] = append(intCols[i], v)
		}
	}
	return intCols
}

func pt1GetColumns() ([][]string, []string) {
	collumns := [][]string{}
	ops := []string{}

	// pt 1
	util.Read(func(line string) {
		parts := strings.Fields(line)
		for len(collumns) < len(parts) {
			collumns = append(collumns, []string{})
		}
		for p := range parts {
			collumns[p] = append(collumns[p], parts[p])
		}
	}, "day6/input.txt")

	for i, col := range collumns {
		ops = append(ops, col[len(col)-1])
		collumns[i] = col[:len(col)-1]
	}

	return collumns, ops
}

func pt2GetColumns() ([][]string, []string) {
	grid := util.ReadMatrix("day6/input.txt")
	columns := make([]string, len(grid[0]))

	for y := range grid[:len(grid)-1] {
		for x := range grid[y] {
			columns[x] += string(grid[y][x])
		}
	}

	columns2 := make([][]string, len(grid[0]))
	for i := range grid {
		columns2[i] = make([]string, len(grid))
	}
	currentY := 0
	for x := range columns {
		if strings.TrimSpace(columns[x]) == "" {
			currentY += 1
			continue
		}
		columns2[currentY] = append(columns2[currentY], columns[x])
	}

	ops := []string{}
	for _, c := range grid[len(grid)-1] {
		if c != ' ' {
			ops = append(ops, string(c))
		}
	}
	return columns2, ops
}

func main() {
	columns, ops := pt1GetColumns()
	pt1Columns := toIntColumns(columns)
	fmt.Println("pt1: ", doTheMath(pt1Columns, ops))

	pt2ColumnsStr, ops := pt2GetColumns()
	pt2Columns := toIntColumns(pt2ColumnsStr)

	fmt.Println("pt2: ", doTheMath(pt2Columns, ops))
}
