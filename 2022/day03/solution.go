package day03

import (
	"fmt"
	"strings"

	"github.com/zvandehy/adventofcode/2022/reader"
)

func Day3() {
	input := reader.Read("../day03/input.txt")

	totalPriority := 0
	commonItemsPriority := 0
	commonItems := make(map[rune]int)
	for i, line := range strings.Split(input, "\n") {
		// by doing part 1 and part 2 together, we check if the index is part of a new group (i%3)
		switch i % 3 {
		case 0:
			commonItems = make(map[rune]int)
			for _, item := range line {
				commonItems[item] = 1
			}
		case 1:
			for _, item := range line {
				if _, ok := commonItems[item]; ok {
					commonItems[item] = 2
				}
			}
		case 2:
			for _, item := range line {
				if x, ok := commonItems[item]; ok && x == 2 {
					commonItems[item] = 3
					commonItemsPriority += Priority(string(item))
					break
				}
			}
		}
		compartment1 := line[:len(line)/2]
		compartment2 := line[len(line)/2:]
		firstHalfContains := make(map[rune]interface{}, len(compartment1))
		for _, char := range compartment1 {
			firstHalfContains[char] = nil
		}
		for _, char := range compartment2 {
			if _, ok := firstHalfContains[char]; ok {
				totalPriority += Priority(string(char))
				break
			}
		}

	}
	fmt.Println("Part 1: ", totalPriority)
	fmt.Println("Part 2: ", commonItemsPriority)
}
