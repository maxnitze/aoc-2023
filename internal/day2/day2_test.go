package day2

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	exampleInputFilePath := "data/example_part1.txt"
	exampleBagLoad := &BagLoad{Red: 12, Green: 13, Blue: 14}
	exampleInputExpectedValue := 8

	content, err := os.ReadFile(exampleInputFilePath)
	if err != nil {
		t.Fatalf("failed to read file at path '%s'", exampleInputFilePath)
	}

	result, err := part1(string(content), exampleBagLoad)
	if err != nil {
		t.Fatalf(`%v`, err)
	}

	if result.Value != exampleInputExpectedValue {
		t.Fatalf(`result for example data '%d' does not match expected value '%d'`, result.Value, exampleInputExpectedValue)
	}
}
