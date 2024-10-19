package twoNumberSum

import "sort"

// TwoNumberSumV1 Complexity: O(n^2) time | O(1) space
func TwoNumberSumV1(array []int, target int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == target {
				return []int{array[i], array[j]}
			}
		}
	}

	return []int{}
}

// TwoNumberSumV2 Complexity: O(n) time | O(n) space
func TwoNumberSumV2(array []int, target int) []int {
	nums := map[int]bool{}

	for _, num := range array {
		potentialMatch := target - num
		if _, found := nums[potentialMatch]; found {
			return []int{potentialMatch, num}
		}
		nums[num] = true
	}

	return []int{}
}

// TwoNumberSumV3 Complexity: O(n log(n)) time | O(1) space
func TwoNumberSumV3(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum < target {
			left += 1
		} else {
			right -= 1
		}
	}

	return []int{}
}
