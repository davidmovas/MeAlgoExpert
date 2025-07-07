package RandomNote_383

func v1(randomNote, magazine string) bool {
	magazineMap := make(map[int32]int)
	for _, let := range magazine {
		magazineMap[let]++
	}

	for _, note := range randomNote {
		_, exists := magazineMap[note]
		if !exists {
			return false
		}
		magazineMap[note]--
		if magazineMap[note] == 0 {
			delete(magazineMap, note)
		}
	}

	return true
}

func v2(randomNote, magazine string) bool {
	lettersCount := [26]int{}

	for _, ch := range magazine {
		lettersCount[ch-'a']++
	}

	for _, ch := range randomNote {
		lettersCount[ch-'a']--
		if lettersCount[ch-'a'] < 0 {
			return false
		}
	}

	return true
}
