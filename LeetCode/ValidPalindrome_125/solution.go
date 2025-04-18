package ValidPalindrome_125

import (
	"strings"
	"unicode"
)

func v1(s string) bool {
	s = strings.ToLower(onlyLetters(s))

	if len(s) == 1 {
		return true
	}

	var left, right = 0, len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func onlyLetters(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
