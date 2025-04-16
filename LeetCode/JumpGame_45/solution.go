package JumpGame_55

// NOT MY SOLUTION
func v1(nums []int) int {
	maxReachable, minJumps, currentEnd, length := 0, 0, 0, len(nums)

	for i := 0; i < length-1; i++ {
		if i+nums[i] > maxReachable {
			maxReachable = i + nums[i]
		}

		if i == currentEnd {
			minJumps++
			currentEnd = maxReachable
		}
	}

	return minJumps
}
