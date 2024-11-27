package algox

type Number interface {
	int | float32 | float64
}

func ArraySum[T Number](array []T) T {
	var result T
	for _, value := range array {
		result += value
	}
	return result
}
