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

func Sort[T Number](array []T, sortFn func(a, b T) bool) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if sortFn(array[j], array[i]) {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
