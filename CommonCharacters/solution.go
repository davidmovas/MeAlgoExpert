package CommonCharacters

func CommonCharacters(sets []string) (result []string) {
	var set = map[rune]int{}
	var length = len(sets) - 1

	for _, ss := range sets {
		for _, s := range ss {
			if _, ok := set[s]; !ok {
				set[s]++
			}
		}
	}

	for key, val := range set {
		if val == length {
			result = append(result, string(key))
		}
	}

	return result
}
