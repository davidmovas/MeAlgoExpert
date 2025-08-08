package ShuffleTheArray_1470

func v1(nums []int, n int) []int {
	var ans []int
	for i := 0; i < n; i++ {
		ans = append(ans, nums[i], nums[i+n])
	}
	return ans
}
