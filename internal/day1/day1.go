package day1

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

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
	firstDigitRe := regexp.MustCompile("^[^\\d]*(\\d).*$")
	lastDigitRe := regexp.MustCompile("^.*(\\d)[^\\d]*$")

	fullCalibrationValue := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		firstDigitMatch := firstDigitRe.FindStringSubmatch(line)
		if len(firstDigitMatch) != 2 {
			return nil, errors.New(fmt.Sprintf("failed to parse first digit from line '%s'", line))
		}
		lastDigitMatch := lastDigitRe.FindStringSubmatch(line)
		if len(lastDigitMatch) != 2 {
			return nil, errors.New(fmt.Sprintf("failed to parse last digit from line '%s'", line))
		}

		calibrationValue, err := strconv.Atoi(fmt.Sprintf("%s%s", firstDigitMatch[1], lastDigitMatch[1]))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to parse calibration value from line '%s'", line))
		}
		fullCalibrationValue += calibrationValue
	}

	message := fmt.Sprintf("The calibration value is '%d'!", fullCalibrationValue)
	return &common.PartResult{Message: message, Value: fullCalibrationValue}, nil
}

func part2(input string) (*common.PartResult, error) {
	return &common.PartResult{Message: "Not Implemented Yet", Value: -1}, nil
}
