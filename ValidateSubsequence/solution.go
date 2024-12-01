package ValidateSubsequence

func ValidateSubsequenceV1(array []int, sequence []int) bool {
	index := 0
	length := len(sequence) - 1
	target := sequence[index]

	for _, num := range array {
		if num == target {
			index++
			target = sequence[index]
			if index == length {
				return true
			}
		}
	}

	return false
}
