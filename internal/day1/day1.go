package day1

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/maxnitze/aoc-2023/internal/common"
)

var firstDigitRe = regexp.MustCompile("^[^\\d]*(\\d).*$")
var lastDigitRe = regexp.MustCompile("^.*(\\d)[^\\d]*$")
var firstSpelledOutDigitRe = regexp.MustCompile("^(.*?)(one|two|three|four|five|six|seven|eight|nine)(.*)$")
var lastSpelledOutDigitRe = regexp.MustCompile("^(.*)(one|two|three|four|five|six|seven|eight|nine)(.*?)$")

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
	return calibrationValueFromLines(input, false)
}

func part2(input string) (*common.PartResult, error) {
	return calibrationValueFromLines(input, true)
}

func calibrationValueFromLines(input string, replaceSpelledOutDigits bool) (*common.PartResult, error) {
	fullCalibrationValue := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var firstDigitLine, lastDigitLine string
		if replaceSpelledOutDigits {
			firstDigitLine = replaceSpelledOutDigit(line, firstSpelledOutDigitRe)
			lastDigitLine = replaceSpelledOutDigit(line, lastSpelledOutDigitRe)
		} else {
			firstDigitLine = line
			lastDigitLine = line
		}

		firstDigitMatch := firstDigitRe.FindStringSubmatch(firstDigitLine)
		if len(firstDigitMatch) != 2 {
			return nil, errors.New(fmt.Sprintf("failed to parse first digit from line '%s'", line))
		}
		lastDigitMatch := lastDigitRe.FindStringSubmatch(lastDigitLine)
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

var textToNumber = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func replaceSpelledOutDigit(line string, regex *regexp.Regexp) string {
	spelledOutDigitMatch := regex.FindStringSubmatch(line)
	if len(spelledOutDigitMatch) == 4 {
		return fmt.Sprintf("%s%d%s", spelledOutDigitMatch[1], textToNumber[spelledOutDigitMatch[2]], spelledOutDigitMatch[3])
	} else {
		return line
	}
}
