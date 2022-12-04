package day04

import (
	"fmt"
	"strings"

	"github.com/zvandehy/adventofcode/2022/reader"
)

func Day4() {
	input := reader.Read("../day04/input.txt")

	fullyContained := 0
	partiallyContained := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		pairs := strings.Split(line, ",")
		a := NewCleaningRange(pairs[0])
		b := NewCleaningRange(pairs[1])
		if FullyContains(a, b) {
			fullyContained++
		}
		if PartiallyContains(a, b) {
			partiallyContained++
		}
	}
	fmt.Println("Part 1:", fullyContained)
	fmt.Println("Part 2:", partiallyContained)
}
