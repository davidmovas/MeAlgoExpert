package Semoedilap

func Semoedilap(words []string) [][]string {
	var result [][]string

	for idx, word := range words {
		for i := idx + 1; i < len(words); i++ {
			if hasPair(word, words[i]) {
				result = append(result, []string{word, words[i]})
			}
		}
	}

	return result
}

func hasPair(origin, potential string) bool {
	if len(origin) != len(potential) {
		return false
	}

	var head, tail = 0, len(origin) - 1

	for head < len(origin) {
		if origin[head] != potential[tail] {
			return false
		}
		head++
		tail--
	}

	return true
}
