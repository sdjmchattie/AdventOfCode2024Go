package main

import (
	"aoc2024/solutions/day01"
	"aoc2024/solutions/day02"
	"fmt"
	"os"
	"strings"
)

type Solver interface {
	Part1([]string) string
	Part2([]string) string
}

var solvers = map[string]Solver{
	"1": day01.Solver{},
	"2": day02.Solver{},
}

func main() {
	var dayNums []string

	if len(os.Args) < 2 {
		for key := range solvers {
			dayNums = append(dayNums, key)
		}
	} else {
		dayNums = strings.Split(os.Args[1], ",")
	}

	for _, dayNum := range dayNums {
		solver, ok := solvers[dayNum]
		if !ok {
			panic("No solver found for day " + dayNum)
		}

		example_input, actual_input := readInputs(dayNum)


		fmt.Println("Day " + dayNum + ":")

		if len(example_input) > 0 {
			fmt.Println("  Example input:")
			fmt.Println("    Part 1:", solver.Part1(example_input))
			fmt.Println("    Part 2:", solver.Part2(example_input))
		}

		fmt.Println("  Actual input:")
		fmt.Println("    Part 1: ", solver.Part1(actual_input))
		fmt.Println("    Part 2: ", solver.Part2(actual_input))
		fmt.Println()
	}
}

func readInputs(dayNum string) ([]string, []string) {
	dayName := fmt.Sprintf("day%02s", dayNum)
	example_input := readLinesFromFile("./inputs/" + dayName + "-example.txt")
	actual_input := readLinesFromFile("./inputs/" + dayName + ".txt")

	return example_input, actual_input
}

func readLinesFromFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		return []string{}
	}

	lines := strings.Split(string(content), "\n")
	return lines[:len(lines)-1] // Last element is a blank line
}
