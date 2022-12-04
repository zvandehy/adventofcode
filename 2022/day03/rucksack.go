package day03

import "strings"

func Priority(item string) int {
	priorities := "abcdefghijklmnopqrstuvwxyz"
	lowercase := strings.Index(priorities, item)
	if lowercase == -1 {
		return 27 + strings.Index(strings.ToUpper(priorities), item)
	}
	return lowercase + 1
}

func PriorityFromCharCode(charCode int) int {
	if charCode >= 65 && charCode <= 90 {
		return charCode - 64 + 26
	}
	return charCode - 96
}
