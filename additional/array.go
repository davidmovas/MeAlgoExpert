package algox

func ArraySum[T int | float64 | float32](array []T) T {
	var result T
	for _, value := range array {
		result += value
	}
	return result
}
