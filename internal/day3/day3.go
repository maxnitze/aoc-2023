package day3

import (
	"errors"

	"github.com/maxnitze/aoc-2023/internal/common"
)

func Solve(input string) (*common.Result, error) {
	if input == "" {
		return nil, errors.New("no input given")
	}

	part1Result, err := part1(input)
	if err != nil {
		return nil, err
	}
	part2Result, err := part2(input)
	if err != nil {
		return nil, err
	}

	return &common.Result{Part1: *part1Result, Part2: *part2Result}, nil
}

func part1(input string) (*common.PartResult, error) {
	return &common.PartResult{Message: "NOT IMPLEMENTED YET", Value: -1}, nil
}

func part2(input string) (*common.PartResult, error) {
	return &common.PartResult{Message: "NOT IMPLEMENTED YET", Value: -1}, nil
}
