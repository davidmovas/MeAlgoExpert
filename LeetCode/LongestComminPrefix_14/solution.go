package LongestComminPrefix_14

import (
	"strings"
)

func v1(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	var builder strings.Builder
	var shortestLen = len(strs[0])
	var shortestIdx int

	for i, str := range strs {
		if shortestLen > len(str) {
			shortestIdx = i
		}
	}

	shortest := strs[shortestIdx]
	arrays := append(strs[:shortestIdx], strs[shortestIdx+1:]...)

	for i := 0; i < len(shortest); i++ {
		same := false

		for _, str := range arrays {
			if str[i] == shortest[i] {
				same = true
			} else {
				return builder.String()
			}
		}

		if same {
			builder.WriteRune(rune(shortest[i]))
			same = false
		}
	}

	return builder.String()
}
