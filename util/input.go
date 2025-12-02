package util

import (
	"bufio"
	"os"
)

type LineParser func(string)

func Read(parser LineParser, filename string) {
	inputFile, error := os.Open("day1/input.txt")
	if error != nil {
		panic(error)
	}

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		parser(line)
	}
}
