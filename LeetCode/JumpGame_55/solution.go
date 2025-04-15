package JumpGame_55

// TODO: NOT MY SOLUTION
func v1(nums []int) bool {
	maxReachable := 0

	for idx := 0; idx < len(nums); idx++ {
		// check if current index is reachable or not
		if idx > maxReachable {
			return false
		}
		// current index is reachable so calculate max index we can jump to
		if idx+nums[idx] > maxReachable {
			maxReachable = idx + nums[idx]
		}
	}
	return true
}
