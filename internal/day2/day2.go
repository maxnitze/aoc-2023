package day2

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/maxnitze/aoc-2023/internal/common"
)

type BagLoad struct {
	Red   int
	Green int
	Blue  int
}

func (b BagLoad) forName(name string) (int, error) {
	if name == "red" {
		return b.Red, nil
	} else if name == "green" {
		return b.Green, nil
	} else if name == "blue" {
		return b.Blue, nil
	} else {
		return -1, errors.New(fmt.Sprintf("unknown name '%s'", name))
	}
}

func Solve(input string, bagLoad *BagLoad) (*common.Result, error) {
	if input == "" {
		return nil, errors.New("no input given")
	}

	part1Result, err := part1(input, bagLoad)
	if err != nil {
		return nil, err
	}
	part2Result, err := part2(input)
	if err != nil {
		return nil, err
	}

	return &common.Result{Part1: *part1Result, Part2: *part2Result}, nil
}

var gameRe = regexp.MustCompile("^Game (\\d+): (.+)$")
var cubeRe = regexp.MustCompile("^(\\d+) (red|green|blue)$")

func part1(input string, bagLoad *BagLoad) (*common.PartResult, error) {
	gameIdSum := 0

game_loop:
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		gameMatch := gameRe.FindStringSubmatch(line)
		if len(gameMatch) != 3 {
			return nil, errors.New(fmt.Sprintf("cannot parse game from line '%s'", line))
		}

		for _, reveal := range strings.Split(gameMatch[2], ";") {
			for _, c := range strings.Split(strings.TrimSpace(reveal), ",") {
				cubeMatch := cubeRe.FindStringSubmatch(strings.TrimSpace(c))
				if len(cubeMatch) != 3 {
					return nil, errors.New(fmt.Sprintf("failed to parse cube from part '%s' of line '%s'", strings.TrimSpace(c), line))
				}

				cubeValue, err := strconv.Atoi(cubeMatch[1])
				if err != nil {
					return nil, errors.New(fmt.Sprintf("failed to parse number of cubes from part '%s' from line '%s'", strings.TrimSpace(c), line))
				}

				bagLoadValueForName, err := bagLoad.forName(cubeMatch[2])
				if err != nil {
					return nil, errors.New(fmt.Sprintf("could not find bag load value for name '%s' from line '%s'", cubeMatch[2], line))
				}

				if cubeValue > bagLoadValueForName {
					continue game_loop
				}
			}
		}

		gameId, err := strconv.Atoi(gameMatch[1])
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to parse game id from line '%s'", line))
		}

		gameIdSum += gameId
	}

	return &common.PartResult{Message: fmt.Sprintf("the sum of the ids of all possible games is '%d'", gameIdSum), Value: gameIdSum}, nil
}

func part2(input string) (*common.PartResult, error) {
	gamePowerSum := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		gameMatch := gameRe.FindStringSubmatch(line)
		if len(gameMatch) != 3 {
			return nil, errors.New(fmt.Sprintf("cannot parse game from line '%s'", line))
		}

		maxBagLoad := BagLoad{Red: 0, Green: 0, Blue: 0}
		for _, reveal := range strings.Split(gameMatch[2], ";") {
			for _, c := range strings.Split(strings.TrimSpace(reveal), ",") {
				cubeMatch := cubeRe.FindStringSubmatch(strings.TrimSpace(c))
				if len(cubeMatch) != 3 {
					return nil, errors.New(fmt.Sprintf("failed to parse cube from part '%s' of line '%s'", strings.TrimSpace(c), line))
				}

				cubeValue, err := strconv.Atoi(cubeMatch[1])
				if err != nil {
					return nil, errors.New(fmt.Sprintf("failed to parse number of cubes from part '%s' from line '%s'", strings.TrimSpace(c), line))
				}

				maxCubeValue, err := maxBagLoad.forName(cubeMatch[2])
				if err != nil {
					return nil, err
				}

				if maxCubeValue < cubeValue {
					if cubeMatch[2] == "red" {
						maxBagLoad.Red = cubeValue
					} else if cubeMatch[2] == "green" {
						maxBagLoad.Green = cubeValue
					} else if cubeMatch[2] == "blue" {
						maxBagLoad.Blue = cubeValue
					} else {
						return nil, errors.New(fmt.Sprintf("unknown cube color '%s' in line '%s'", cubeMatch[2], line))
					}
				}
			}
		}

		if maxBagLoad.Red <= 0 || maxBagLoad.Green <= 0 || maxBagLoad.Blue <= 0 {
			return nil, errors.New(fmt.Sprintf("failed to get max bag load values for all three colors from line '%s'", line))
		}

		gamePowerSum += (maxBagLoad.Red * maxBagLoad.Green * maxBagLoad.Blue)
	}

	return &common.PartResult{Message: fmt.Sprintf("the sum of the powers of all games is '%d'", gamePowerSum), Value: gamePowerSum}, nil
}
