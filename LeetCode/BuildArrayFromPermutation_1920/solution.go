package BuildArrayFromPermutation_1920

import "runtime/debug"

func v1(nums []int) []int {
	var ans = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}

	return ans
}

func init() {
	debug.SetMemoryLimit(1)
}
