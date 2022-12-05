package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zvandehy/adventofcode/2022/reader"
)

func Day5() {
	input := reader.Read("../day05/input.txt")
	crates := make(map[int][]string)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		if line[1] == '1' {
			break
		}
		crateNum := 0
		for i := 1; i < len(line); i += 4 {
			crateNum++
			// fmt.Print(string(line[i]))
			if line[i] == ' ' {
				continue
			}
			if _, ok := crates[crateNum]; !ok {
				crates[crateNum] = []string{string(line[i])}

			} else {
				crates[crateNum] = append([]string{string(line[i])}, crates[crateNum]...)
			}
		}
		// fmt.Println()
	}

	stacks := make(map[int]*Stack)

	for i := 1; i <= len(crates); i++ {
		stacks[i] = New()
		for _, crate := range crates[i] {
			stacks[i].Push(crate)
		}
	}

	for _, line := range strings.Split(input, "\n") {
		if line == "" || line[0] != 'm' {
			continue
		}
		x, _ := strconv.Atoi(strings.Split(line, " ")[1])
		from, _ := strconv.Atoi(strings.Split(line, " ")[3])
		to, _ := strconv.Atoi(strings.Split(line, " ")[5])
		// fmt.Println(x, from, to)
		for i := 0; i < x; i++ {
			stacks[to].Push(stacks[from].Pop())
		}

		// fmt.Println("move", x, crates[from], crates[from][len(crates[from])-x:], "from", from, "to", to)
		crates[to] = append(crates[to], crates[from][len(crates[from])-x:]...)
		crates[from] = crates[from][:len(crates[from])-x]
		// fmt.Println(crates, "\n")
	}

	fmt.Print("Part 1: ")
	for i := 1; i <= len(stacks); i++ {
		fmt.Print(stacks[i].Peek())
	}
	fmt.Println()
	fmt.Print("Part 2: ")
	for i := 1; i <= len(crates); i++ {
		n := len(crates[i])
		if n > 0 {
			fmt.Print(crates[i][n-1])
		}
	}
	fmt.Println()
}
