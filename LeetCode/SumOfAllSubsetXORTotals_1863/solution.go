package SumOfAllSubsetXORTotals_1863

func v1(nums []int) int {
	total := 0
	n := len(nums)

	for mask := 0; mask < (1 << n); mask++ {
		currentXOR := 0
		for i := 0; i < n; i++ {
			if (mask & (1 << i)) != 0 {
				currentXOR ^= nums[i]
			}
		}
		total += currentXOR
	}

	return total
}
