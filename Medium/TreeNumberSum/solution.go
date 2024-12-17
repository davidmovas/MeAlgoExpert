package TreeNumberSum

import "sort"

func TreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	var triplets [][]int

	for i := 0; i < len(array)-2; i++ {
		left, right := i+1, len(array)-1
		for left < right {
			currentSum := array[i] + array[left] + array[right]
			if currentSum == target {
				triplets = append(triplets, []int{array[i], array[left], array[right]})
				left++
				right--
			} else if currentSum < target {
				left++
			} else if currentSum > target {
				right--
			}
		}
	}
	return triplets
}
