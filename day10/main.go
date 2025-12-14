package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"AOC_2025/util"
)

func bfs(target int, buttons []int) int {
	currentValues := []int{0}
	stepCount := 0
	for {
		nextValues := []int{}
		stepCount += 1
		for i := range len(currentValues) {
			for _, button := range buttons {
				new := currentValues[i] ^ button
				nextValues = append(nextValues, new)
				if new == target {
					return stepCount
				}
			}
		}
		currentValues = nextValues
	}
}

type buttonCombo struct {
	buttonIndexs []int
	maxCount     int
}

func exceedsMaxCombinedCount(presses []int, maxCombinedCount []buttonCombo) bool {
	for _, combo := range maxCombinedCount {
		sum := 0
		for _, btnIdx := range combo.buttonIndexs {
			sum += presses[btnIdx]
		}
		if sum > combo.maxCount {
			return true
		}
	}
	return false
}

func checkCombo(targetJolts []int, buttons [][]int, presses []int) bool {
	joltage := make([]int, len(targetJolts))
	for i, p := range presses {
		bttn := buttons[i]
		for _, j := range bttn {
			joltage[j] += p
		}
	}

	for i := range targetJolts {
		if targetJolts[i] != joltage[i] {
			return false
		}
	}
	return true
}

func tryCombinationsWith(targetJolts []int, buttons [][]int, maxBttnPresses map[int]int, maxCombined []buttonCombo) int {
	currentPresses := [][]int{make([]int, len(buttons))}
	stepCount := 0
	for {
		nextValues := [][]int{}
		stepCount += 1
		for i := range len(currentPresses) {
		bttnLoop:
			for bi := range buttons {
				newPress := slices.Clone(currentPresses[i])
				newPress[bi] += 1
				if maxBttnPresses[bi] < newPress[bi] || exceedsMaxCombinedCount(newPress, maxCombined) {
					continue bttnLoop
				}
				nextValues = append(nextValues, newPress)
				if checkCombo(targetJolts, buttons, newPress) {
					return stepCount
				}
			}
		}
		currentPresses = nextValues
	}
}

func pt2bfs(targetJolts []int, buttons [][]int) int {
	maxBttnPresses := map[int]int{}
	for i, joltage := range targetJolts {
		for bttnI, bttn := range buttons {
			if slices.Contains(bttn, i) {
				_, ok := maxBttnPresses[bttnI]
				if !ok {
					maxBttnPresses[bttnI] = joltage
				} else if maxBttnPresses[bttnI] > joltage {
					maxBttnPresses[bttnI] = joltage
				}
			}
		}
	}

	maxCombinedBtnPresses := []buttonCombo{}
	for i, joltage := range targetJolts {
		bttnsWithInfluence := []int{}
		for bttnI, bttn := range buttons {
			if slices.Contains(bttn, i) {
				bttnsWithInfluence = append(bttnsWithInfluence, bttnI)
			}
		}
		maxCombinedBtnPresses = append(maxCombinedBtnPresses, buttonCombo{
			buttonIndexs: bttnsWithInfluence,
			maxCount:     joltage,
		})
	}

	steps := tryCombinationsWith(targetJolts, buttons, maxBttnPresses, maxCombinedBtnPresses)
	fmt.Println(steps)
	return steps
}

func main() {
	stepsCountPt1 := 0
	stepsCountPt2 := 0
	util.Read(func(line string) {
		s1 := strings.Split(line, "]")
		lightsTarget := s1[0]
		s2 := strings.Split(s1[1], "{")
		buttons := s2[0]
		joltsStr := s2[1]

		targetLightInt := 0
		for i, c := range lightsTarget[1:] {
			if c == '#' {
				targetLightInt = targetLightInt | (1 << i)
			}
		}

		coords := strings.Split(strings.TrimSpace(buttons), " ")
		buttonI := []int{}
		buttonJ := [][]int{}
		for cidx, c := range coords {
			r := strings.Split(c[1:len(c)-1], ",")
			buttonJ = append(buttonJ, []int{})
			buttonSet := 0
			for _, cs := range r {
				ci, err := strconv.Atoi(cs)
				if err != nil {
					panic(err)
				}
				buttonSet += 1 << ci
				buttonJ[cidx] = append(buttonJ[cidx], ci)
			}
			buttonI = append(buttonI, buttonSet)
		}

		joltPS := strings.Split(joltsStr[:len(joltsStr)-1], ",")
		jolts := []int{}
		for _, js := range joltPS {
			i, e := strconv.Atoi(js)
			if e != nil {
				panic(e)
			}
			jolts = append(jolts, i)
		}

		stepsCountPt1 += bfs(targetLightInt, buttonI)
		stepsCountPt2 += pt2bfs(jolts, buttonJ)
	}, "day10/input.txt")

	fmt.Println("pt1: ", stepsCountPt1)
	fmt.Println("pt2: ", stepsCountPt2)
}
