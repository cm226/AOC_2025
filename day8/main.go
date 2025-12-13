package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"AOC_2025/util"
)

type Point struct {
	X int
	Y int
	Z int
}

type Connection struct {
	p1       Point
	p2       Point
	distance int
}

type Connections []Connection

func (r Connections) Len() int { return len(r) }
func (r Connections) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r Connections) Less(i, j int) bool { return r[i].distance < r[j].distance }

func (p Point) DistSqr(o Point) int {
	x := (p.X - o.X)
	y := (p.Y - o.Y)
	z := (p.Z - o.Z)
	return x*x + y*y + z*z
}

func connectClosest(sortedDistances Connections, graph [][]Point, i int) [][]Point {
	next := sortedDistances[i]
	circetP1 := slices.IndexFunc(graph, func(circet []Point) bool {
		return slices.IndexFunc(circet, func(p Point) bool { return p == next.p1 }) != -1
	})

	circetP2 := slices.IndexFunc(graph, func(circet []Point) bool {
		return slices.IndexFunc(circet, func(p Point) bool { return p == next.p2 }) != -1
	})

	if circetP1 == -1 || circetP2 == -1 {
		panic("failed to find circet")
	}
	if circetP1 == circetP2 {
		return graph
	}

	graph[circetP1] = append(graph[circetP1], graph[circetP2]...)
	graph[circetP2] = []Point{}
	return graph
}

func main() {
	junctions := []Point{}

	util.Read(func(line string) {
		values := strings.Split(line, ",")

		X, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		Y, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		Z, err := strconv.Atoi(values[2])
		if err != nil {
			panic(err)
		}

		point := Point{X: X, Y: Y, Z: Z}
		junctions = append(junctions, point)
	}, "day8/input.txt")

	distances := Connections{}
	for i := range junctions[:len(junctions)-1] {
		for _, j := range junctions[i+1:] {
			distances = append(distances, Connection{
				p1:       junctions[i],
				p2:       j,
				distance: junctions[i].DistSqr(j),
			})
		}
	}

	sort.Sort(Connections(distances))
	fmt.Println(distances[0])

	graph := [][]Point{}
	for _, p := range junctions {
		graph = append(graph, []Point{p})
	}

	for i := range 1000 {
		graph = connectClosest(distances, graph, i)
	}

	slices.SortFunc(graph, func(a, b []Point) int {
		return len(b) - len(a)
	})

	fmt.Println("Pt1 : ", len(graph[0])*len(graph[1])*len(graph[2]))

	i := 1000
	for len(graph[1]) != 0 {
		graph = connectClosest(distances, graph, i)
		slices.SortFunc(graph, func(a, b []Point) int {
			return len(b) - len(a)
		})
		i++
	}

	fmt.Println("Pt2: ", distances[i-1].p1.X*distances[i-1].p2.X)
}
