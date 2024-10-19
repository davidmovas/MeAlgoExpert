package main

import (
	algox "algoexpert/additional"
	"fmt"
)

func main() {
	fmt.Printf("Result: %v\n", algox.ArraySum(twoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)) == algox.ArraySum([]int{11, -1}))
}

func twoNumberSum(array []int, target int) []int {
	var result []int

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == target {
				result = append(result, array[i], array[j])
			}
		}
	}

	return result
}
