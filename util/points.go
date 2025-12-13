package util

import (
	"strconv"
	"strings"

	"github.com/s0rg/vec2d"
)

func Up() vec2d.V[int]    { return vec2d.New(0, -1) }
func Down() vec2d.V[int]  { return vec2d.New(0, 1) }
func Left() vec2d.V[int]  { return vec2d.New(-1, 0) }
func Right() vec2d.V[int] { return vec2d.New(1, 0) }

func All() []vec2d.V[int] {
	return []vec2d.V[int]{
		Up(), Down(), Left(), Right(),
		Up().Add(Left()), Up().Add(Right()), Down().Add(Left()), Down().Add(Right()),
	}
}

func FromText(line string) vec2d.V[int] {
	parts := strings.Split(line, ",")
	p1, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	p2, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return vec2d.New(p1, p2)
}
