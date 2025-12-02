package util

import (
	"bufio"
	"os"
	"strings"
)

type LineParser func(string)

func Read(parser LineParser, filename string) {
	inputFile, error := os.Open(filename)
	if error != nil {
		panic(error)
	}

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		parser(line)
	}
}

func ReadCSV(parser LineParser, filename string) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	elements := strings.SplitSeq(string(dat), ",")
	for element := range elements {
		parser(element)
	}
}
