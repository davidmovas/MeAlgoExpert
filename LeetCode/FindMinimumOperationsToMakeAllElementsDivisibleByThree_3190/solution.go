package FindMinimumOperationsToMakeAllElementsDivisibleByThree_3190

func v1(nums []int) int {
	var ops int

	for _, num := range nums {
		if num%3 != 0 {
			ops++
		}
	}

	return ops
}
