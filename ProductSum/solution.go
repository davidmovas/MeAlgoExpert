package ProductSum

import (
	"reflect"
)

func ProductSum(array []any) int {
	return countSum(array, 0, 0)
}

func countSum(arr []any, sum, depth int) int {
	depth++
	for _, el := range arr {
		if reflect.TypeOf(el).Kind() == reflect.Slice {
			sum += countSum((el).([]any), 0, depth)
		}
		if reflect.TypeOf(el).Kind() == reflect.Int {
			sum += (el).(int)
		}
	}

	return sum * depth
}
