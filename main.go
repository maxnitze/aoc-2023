package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/maxnitze/aoc-2023/internal/common"
	"github.com/maxnitze/aoc-2023/internal/day1"
	"github.com/maxnitze/aoc-2023/internal/day2"
	"github.com/maxnitze/aoc-2023/internal/day3"
)

func main() {
	var day int
	flag.IntVar(&day, "day", 0, "the day to run")
	var inputFilePath string
	flag.StringVar(&inputFilePath, "input", "", "the input file")
	flag.Parse()

	if inputFilePath == "" {
		inputFilePath = fmt.Sprintf("internal/day%d/data/input.txt", day)
	}

	var result *common.Result
	var err error
	switch day {
	case 1:
		result, err = day1.Solve(readInput(inputFilePath))
	case 2:
		result, err = day2.Solve(readInput(inputFilePath), &day2.BagLoad{Red: 12, Green: 13, Blue: 14})
	case 3:
		result, err = day3.Solve(readInput(inputFilePath))
	default:
		err = errors.New(fmt.Sprintf("Day %d out of range!", day))
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Part1: %s", result.Part1.Message))
	fmt.Println(fmt.Sprintf("Part2: %s", result.Part2.Message))
}

func readInput(inputFilePath string) string {
	content, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to read file at path '%s'", inputFilePath))
	}
	return string(content)
}
