package day02

// I wish I setup the Plays so that there was a Rock, Paper, Scissors type that worked for both players. This seems like not the best way to represent it.
// Maybe Input1 type and Input2 type, and then a function maps them to the correct play?

type RockPaperScissors struct {
	OpponentPlay string
	YourPlay     string
	RequiredEnd  string
}

func (rps RockPaperScissors) GetYourPlayFromRequiredEnd() YourPlay {
	switch rps.RequiredEnd {
	case string(Draw):
		switch rps.OpponentPlay {
		case string(OpponentRock):
			return YourRock
		case string(OpponentPaper):
			return YourPaper
		case string(OpponentScissors):
			return YourScissors
		}
	case string(Win):
		switch rps.OpponentPlay {
		case string(OpponentRock):
			return YourPaper
		case string(OpponentPaper):
			return YourScissors
		case string(OpponentScissors):
			return YourRock
		}
	case string(Lose):
		switch rps.OpponentPlay {
		case string(OpponentRock):
			return YourScissors
		case string(OpponentPaper):
			return YourRock
		case string(OpponentScissors):
			return YourPaper
		}
	}
	return YourRock
}

func (rps RockPaperScissors) Score() int {
	switch rps.YourPlay {
	case string(YourRock):
		return rps.Outcome() + 1
	case string(YourPaper):
		return rps.Outcome() + 2
	case string(YourScissors):
		return rps.Outcome() + 3
	}
	return 0
}

func (rps RockPaperScissors) Outcome() int {
	switch rps.OpponentPlay {
	case string(OpponentRock):
		switch rps.YourPlay {
		case string(YourRock):
			return 3
		case string(YourPaper):
			return 6
		case string(YourScissors):
			return 0
		}
	case string(OpponentPaper):
		switch rps.YourPlay {
		case string(YourRock):
			return 0
		case string(YourPaper):
			return 3
		case string(YourScissors):
			return 6
		}
	case string(OpponentScissors):
		switch rps.YourPlay {
		case string(YourRock):
			return 6
		case string(YourPaper):
			return 0
		case string(YourScissors):
			return 3
		}
	}
	return 0
}

type OpponentPlay string
type YourPlay string
type RequiredEnd string

const (
	OpponentRock     OpponentPlay = "A"
	OpponentPaper    OpponentPlay = "B"
	OpponentScissors OpponentPlay = "C"
)

const (
	YourRock     YourPlay = "X"
	YourPaper    YourPlay = "Y"
	YourScissors YourPlay = "Z"
)

const (
	Lose RequiredEnd = "X"
	Draw RequiredEnd = "Y"
	Win  RequiredEnd = "Z"
)
