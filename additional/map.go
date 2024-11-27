package algox

import "fmt"

func ArrayToMap[T comparable, V any](array []T, values []V) (map[T]V, error) {
	if len(array) != len(values) {
		return make(map[T]V), fmt.Errorf("arrays of keys and values not equel to each other")
	}

	result := make(map[T]V)

	for i := 0; i < len(array)-1; i++ {
		if _, ok := result[array[i]]; ok {
			continue
		}

		result[array[i]] = values[i]
	}

	return result, nil
}
