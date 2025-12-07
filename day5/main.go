package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"AOC_2025/util"
)

func pt1(ingreds []int, ranges []util.IntRange) int {
	freshC := 0
ingredLoop:
	for _, ingred := range ingreds {
		for _, fresh := range ranges {
			if fresh.InsideInclusive(ingred) {
				freshC += 1
				continue ingredLoop
			}
		}
	}
	return freshC
}

func pt2(ranges []util.IntRange) int {
	sort.Sort(util.ByStart(ranges))

	nonOverlappingRanges := []util.IntRange{}
	nonOverlappingRanges = append(nonOverlappingRanges, ranges[0])
rngLoop:
	for _, rng := range ranges {
		for i, nonOverlap := range nonOverlappingRanges {
			if nonOverlap.InsideInclusive(rng.Start) {
				nonOverlappingRanges[i].End = int(math.Max(float64(nonOverlap.End), float64(rng.End)))
				continue rngLoop
			}
		}
		nonOverlappingRanges = append(nonOverlappingRanges, rng)
	}

	overlappingSum := 0
	for _, rng := range nonOverlappingRanges {
		overlappingSum += rng.End - (rng.Start - 1)
	}
	return overlappingSum
}

func main() {
	readingIngreds := false
	ingreds := []int{}
	ranges := []util.IntRange{}

	util.Read(func(line string) {
		if line == "" {
			readingIngreds = true
			return
		}

		if readingIngreds {
			i, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			ingreds = append(ingreds, i)
		} else {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				panic("parse error")
			}
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			ranges = append(ranges, util.IntRange{Start: start, End: end})

		}
	}, "day5/input.txt")

	fmt.Println(pt1(ingreds, ranges))
	fmt.Println(pt2(ranges))
}
