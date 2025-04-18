package IsSubsequence_392

import (
	"strings"
)

func v1(s string, t string) bool {
	if s == t {
		return true
	}

	if len(s) == 1 || len(t) == 0 && s != t {
		return false
	}

	str := removeLetters(s, t)

	if idx := strings.Index(s, str); idx == -1 {
		return false
	} else {
		return true
	}
}

func v2(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

func removeLetters(s string, t string) string {
	var builder strings.Builder
	var idx int

	for _, r := range t {
		if idx = strings.IndexRune(s, r); idx != -1 {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
