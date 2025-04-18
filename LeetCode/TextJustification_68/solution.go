package TextJustification_68

import "strings"

func v1(words []string, maxWidth int) []string {
	var result []string
	var current []string
	letters := 0

	for _, word := range words {
		if len(word)+len(current)+letters > maxWidth {
			for i := 0; i < maxWidth-letters; i++ {
				current[i%max(1, len(current)-1)] += " "
			}
			result = append(result, strings.Join(current, ""))
			current = current[:0]
			letters = 0
		}
		current = append(current, word)
		letters += len(word)
	}

	lastLine := strings.Join(current, " ")
	for len(lastLine) < maxWidth {
		lastLine += " "
	}
	result = append(result, lastLine)

	return result
}
