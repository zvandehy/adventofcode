package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/zvandehy/adventofcode/2022/reader"
)

func Day1() {
	fmt.Println("Day 1")
	input := reader.Read("../day01/input.txt")
	elves := []Elf{}
	tempElf := Elf{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			elves = append(elves, Elf{TotalCalories: tempElf.TotalCalories, Items: tempElf.Items})
			tempElf = Elf{}
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		tempElf.TotalCalories += calories
		tempElf.Items = append(tempElf.Items, calories)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].TotalCalories > elves[j].TotalCalories
	})
	maxElf := elves[0]
	topElvesTotal := 0
	for _, elf := range elves[:3] {
		topElvesTotal += elf.TotalCalories
	}
	fmt.Println("Part 1:", maxElf.TotalCalories)
	fmt.Println("Part 2:", topElvesTotal)
}
