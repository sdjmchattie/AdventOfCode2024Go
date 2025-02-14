package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Solver struct{}

var mulRegexp = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
var voidRegexp = regexp.MustCompile(`don't\(\).*?do\(\)`)

func (s Solver) Part1(input []string) string {
	oneLine := strings.Join(input, "")
	return fmt.Sprintf("%d", doMul(oneLine))
}

func (s Solver) Part2(input []string) string {
	oneLine := strings.Join(input, "")
	oneLine = voidRegexp.ReplaceAllString(oneLine, "")
	return fmt.Sprintf("%d", doMul(oneLine))
}

func doMul(input string) int {
	matches := mulRegexp.FindAllStringSubmatch(input, -1)

	total := 0
	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])
		total += left * right
	}

	return total
}
