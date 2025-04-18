package RomanInInteger_13

var romanianMap = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func v1(s string) int {
	var result int
	var runes = []rune(s)
	var current, prev int

	for i := len(runes) - 1; i >= 0; i-- {
		current = romanianMap[runes[i]]

		if current < prev {
			result -= current
		} else {
			result += current
		}

		prev = current
	}

	return result
}
