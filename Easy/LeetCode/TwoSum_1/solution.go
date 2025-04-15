package TwoSum_1

import (
	"runtime"
	"sort"
)

func v1(nums []int, target int) []int {
	var result = make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0], result[1] = i, j
				return result
			}
		}
	}

	return result
}

// nums = [2,7,11,15], target = 9
// nums = [3, 2, 4], target = 6
// sortedNums = [2,3,4]

func v2(nums []int, target int) []int {
	var result = make([]int, 2)

	sort.Ints(nums)

	start, end := 0, len(nums)-1

	for start < end {
		if nums[end] >= target {
			end--
			continue
		}

		if nums[start]+nums[end] == target {
			result[0], result[1] = start, end
			return result
		} else {
			start++
		}
	}

	return result
}

// TODO: NOT MY SOLUTION
func v3(nums []int, target int) []int {
	defer runtime.GC()
	var numsMap = make(map[int]int, len(nums))

	for i, num := range nums {
		complement := target - num
		if index, exist := numsMap[complement]; exist {
			return []int{index, i}
		}

		numsMap[num] = i
	}

	return nil
}
