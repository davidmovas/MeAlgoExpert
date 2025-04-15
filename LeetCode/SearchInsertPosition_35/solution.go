package SearchInsertPosition_35

// 1,3,5,6, target = 2

func v1(nums []int, target int) int {
	for i, num := range nums {
		if num == target || num > target {
			return i
		} else if i == len(nums)-1 {
			return i + 1
		}
	}

	return 0
}
