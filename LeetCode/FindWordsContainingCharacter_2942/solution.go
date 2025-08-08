package FindWordsContainingCharacter_2942

import (
	"bytes"
	"strings"
)

func v1(words []string, x byte) []int {
	var indexes []int
	sub := []byte{x}

	for i, world := range words {
		if bytes.Contains([]byte(world), sub) {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func v2(words []string, x byte) []int {
	var indexes []int

	for i, world := range words {
		if strings.IndexByte(world, x) != -1 {
			indexes = append(indexes, i)
		}
	}

	return indexes
}
