package RotateArray_189

func v1(nums []int, k int) []int {
	l := len(nums) - 1

	for k > 0 {
		var temp = make([]int, l+1)
		temp[0] = nums[l]

		for i := 0; i < l; i++ {
			temp[i+1] = nums[i]
		}

		copy(nums, temp)
		k--
	}

	return nums
}

func v2(nums []int, k int) []int {
	if len(nums) <= 1 {
		return nums
	}

	fStart, fEnd := 0, len(nums)-k
	sStart, sEnd := fEnd, len(nums)-1

	var fPart = nums[fStart:fEnd]
	var sPart = nums[sStart : sEnd+1]

	var result = append(sPart, fPart...)

	copy(nums, result)

	return nums
}

// nums = [1,2,3,4,5,6,7], k = 3
// fStart = 0, fEnd = (7-1-3) = 3
// sStart = 4, sEnd = 6

// Not My
func v3(nums []int, k int) []int {
	if k%len(nums) == 0 {
		return nums
	}
	k = k % len(nums)
	reverse(nums[len(nums)-k : len(nums)])
	reverse(nums[0 : len(nums)-k])
	reverse(nums)
	return nums
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
