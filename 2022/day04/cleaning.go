package day04

import (
	"strconv"
	"strings"
)

type CleaningRange struct {
	From int
	To   int
}

func NewCleaningRange(input string) CleaningRange {
	leftRight := strings.SplitN(input, "-", 2)
	left, err := strconv.Atoi(leftRight[0])
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(leftRight[1])
	if err != nil {
		panic(err)
	}
	return CleaningRange{
		From: left,
		To:   right,
	}
}

func FullyContains(a, b CleaningRange) bool {
	return a.From <= b.From && a.To >= b.To || a.From >= b.From && a.To <= b.To
}

func PartiallyContains(a, b CleaningRange) bool {
	if a.From <= b.From && a.To >= b.From {
		return true
	}
	if b.From <= a.From && b.To >= a.From {
		return true
	}
	return false
}
