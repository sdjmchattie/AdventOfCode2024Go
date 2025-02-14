package day02

import (
	"aoc2024/lib"
	"fmt"
	"math"
	"strings"
)

type Solver struct{}

func (s Solver) Part1(input []string) string {
	reports := parseInput(input)
	safeCount := 0

	for _, report := range reports {
		if isSafe(report, 0) {
			safeCount++
		}
	}

	return fmt.Sprintf("%d", safeCount)
}

func (s Solver) Part2(input []string) string {
	reports := parseInput(input)
	safeCount := 0

	for _, report := range reports {
		if isSafe(report, 1) {
			safeCount++
		}
	}

	return fmt.Sprintf("%d", safeCount)
}

func parseInput(input []string) [][]int {
	lines := [][]int{}
	for _, line := range input {
		fields := strings.Fields(line)
		lines = append(lines, lib.MapInts(fields))
	}

	return lines
}

func isSafe(report []int, tolerance int) bool {
	var sign bool
	remainingTolerance := tolerance
	leftVal := report[0]

	for i := 0; i < len(report) - 1; i++ {
		rightVal := report[i + 1]
		if i == tolerance - remainingTolerance {
			sign = math.Signbit(float64(rightVal - leftVal))
		}

		diff := math.Abs(float64(rightVal - leftVal))
		if math.Signbit(float64(rightVal - leftVal)) != sign || diff < 1 || diff > 3 {
			remainingTolerance--
			if remainingTolerance < 0 {
				return false
			}
		} else {
			leftVal = rightVal
		}
	}

	return true
}
