package ConcatenationOfArray_1929

func v1(nums []int) []int {
	var ans = make([]int, len(nums)*2)
	var n = len(nums)

	for i := 0; i < n; i++ {
		ans[i], ans[i+n] = nums[i], nums[i]
	}

	return ans
}
