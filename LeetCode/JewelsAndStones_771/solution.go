package JewelsAndStones_771

func v1(jewels string, stones string) int {
	var set = make(map[rune]bool, len(jewels))
	for _, j := range jewels {
		set[j] = true
	}

	var total int

	for _, stone := range stones {
		if set[stone] {
			total++
		}
	}

	return total
}
