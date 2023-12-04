package common

type Result struct {
	Part1 PartResult
	Part2 PartResult
}

type PartResult struct {
	Value   int
	Message string
}
