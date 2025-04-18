package FindTheIndexFirstOccurrenceInAString_28

import "strings"

func v1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func v2(haystack string, needle string) int {
	if len(needle) == 0 {
		return -1
	}

	var str, substr = []rune(haystack), []rune(needle)
	n, m := len(str), len(substr)

	for i := 0; i <= n-m; i++ {
		j := 0
		for ; j < m; j++ {
			if str[i+j] != substr[j] {
				break
			}
		}
		if j == m {
			return i
		}
	}

	return -1
}
