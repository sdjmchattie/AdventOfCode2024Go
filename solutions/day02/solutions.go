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
		if isSafe(report) {
			safeCount++
		}
	}

	return fmt.Sprintf("%d", safeCount)
}

func (s Solver) Part2(input []string) string {
	reports := parseInput(input)
	safeCount := 0

	for _, report := range reports {
		for idx := 0; idx < len(report); idx++ {
			truncatedReport := make([]int, len(report) - 1)
			copy(truncatedReport, report[:idx])
			copy(truncatedReport[idx:], report[idx + 1:])
			if isSafe(truncatedReport) {
				safeCount++
				break
			}
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

func isSafe(report []int) bool {
	sign := math.Signbit(float64(report[1] - report[0]))

	for i := 0; i < len(report) - 1; i++ {
		leftVal := report[i]
		rightVal := report[i + 1]

		diff := math.Abs(float64(rightVal - leftVal))
		if math.Signbit(float64(rightVal - leftVal)) != sign || diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}
