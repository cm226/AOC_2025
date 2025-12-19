package main

import (
	"fmt"
	"strings"

	"AOC_2025/util"
)

func dfs(current string, target string, cache map[string]int, machines map[string][]string) int {
	if current == target {
		return 1
	}
	sum := 0
	for _, path := range machines[current] {
		v, ok := cache[path]
		if ok {
			sum += v
			continue
		}
		count := dfs(path, target, cache, machines)
		cache[path] = count
		sum += count

	}
	return sum
}

func p1(machines map[string][]string) {
	start := "you"
	currentPath := map[string]int{}
	fmt.Println(dfs(start, "out", currentPath, machines))
}

func p2(machines map[string][]string) {
	cache := map[string]int{}
	svrfft := dfs("svr", "fft", cache, machines)
	cache = map[string]int{}
	svrdac := dfs("svr", "dac", cache, machines)
	cache = map[string]int{}
	fftdac := dfs("fft", "dac", cache, machines)
	cache = map[string]int{}
	dacfft := dfs("dac", "fft", cache, machines)
	cache = map[string]int{}
	fftout := dfs("fft", "out", cache, machines)
	cache = map[string]int{}
	dacout := dfs("dac", "out", cache, machines)

	svrdacfftout := (svrdac - (svrfft * fftdac)) * dacfft * fftout
	svrfftdacout := (svrfft - (svrdac * dacfft)) * fftdac * dacout

	fmt.Println(svrdacfftout + svrfftdacout)
}

func main() {
	machines := map[string][]string{}
	util.Read(func(line string) {
		parts := strings.Split(line, ":")
		machine := parts[0]
		outputs := strings.Split(strings.TrimSpace(parts[1]), " ")
		machines[machine] = outputs
	}, "day11/input.txt")

	p1(machines)
	p2(machines)
}
