package LengthOfLastWord_58

import (
	"strings"
)

func v1(s string) int {
	var result int
	var runes = []rune(strings.TrimSpace(s))

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == ' ' {
			return result
		}
		result++
	}

	return result
}

func v2(s string) int {
	var result int
	var started bool

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if started {
				break
			}
		} else {
			result++
			started = true
		}
	}

	return result
}
