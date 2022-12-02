package day02

import (
	"fmt"
	"strings"

	"github.com/zvandehy/adventofcode/2022/reader"
)

func Day2() {
	fmt.Println("Day 2")
	input := reader.Read("../day02/input.txt")
	totalScore := 0
	part2Score := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		strategy := strings.Split(line, " ")
		game := RockPaperScissors{
			OpponentPlay: strategy[0],
			YourPlay:     strategy[1],
		}
		totalScore += game.Score()

		game2 := RockPaperScissors{
			OpponentPlay: strategy[0],
			RequiredEnd:  strategy[1],
		}
		game2.YourPlay = string(game2.GetYourPlayFromRequiredEnd())
		part2Score += game2.Score()
	}
	fmt.Println("Part 1: ", totalScore)
	fmt.Println("Part 2: ", part2Score)
}
