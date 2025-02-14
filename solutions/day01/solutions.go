package day01

import (
	"aoc2024/lib"
	"fmt"
	"math"
	"slices"
	"strings"
)

type Solver struct{}

func (s Solver) Part1(input []string) string {
	left, right := parseInput(input)
	slices.Sort(left)
	slices.Sort(right)

	diffs := 0
	for idx := range len(left) {
		diffs += int(math.Abs((float64(right[idx] - left[idx]))))
	}

	return fmt.Sprintf("%d", diffs)
}

func (s Solver) Part2(input []string) string {
	left, right := parseInput(input)
	rightCounts := map[int]int{}
	for _, v := range right {
		rightCounts[v]++
	}

	score := 0
	for _, value := range left {
		score += value * rightCounts[value]
	}

	return fmt.Sprintf("%d", score)
}

func parseInput(input []string) (left []int, right []int) {
	for _, line := range input {
		fields := strings.Fields(line)
		var intVals = lib.MapInts(fields)
		left = append(left, intVals[0])
		right = append(right, intVals[1])
	}

	return left, right
}
