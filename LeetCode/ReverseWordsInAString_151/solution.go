package ReverseWordsInAString_151

import (
	"strings"
)

func v1(s string) string {
	strs := strings.Fields(s)

	var start, end = 0, len(strs) - 1

	for start < end {
		strs[start], strs[end] = strs[end], strs[start]
		start++
		end--
	}

	return strings.Join(strs, " ")
}
