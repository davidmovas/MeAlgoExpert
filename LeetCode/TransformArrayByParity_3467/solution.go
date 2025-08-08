package TransformArrayByParity_3467

import "sort"

func v1(nums []int) []int {
	var ans = make([]int, len(nums))

	for i, num := range nums {
		if num%2 == 0 {
			ans[i] = 0
		} else {
			ans[i] = 1
		}
	}

	sort.Ints(ans)
	return ans
}
