package TwoSum2_167

func v1(numbers []int, target int) []int {
	var head, tail = 0, len(numbers) - 1

	var sum int
	for head < tail {
		sum = numbers[head] + numbers[tail]
		if sum == target {
			return []int{head + 1, tail + 1}
		} else if sum > target {
			tail--
		} else {
			head++
		}
	}

	return []int{}
}
