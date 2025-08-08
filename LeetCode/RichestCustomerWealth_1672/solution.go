package RichestCustomerWealth_1672

func v1(accounts [][]int) int {
	var maximum int

	for _, account := range accounts {
		var total int
		for _, amount := range account {
			total += amount
		}
		if total > maximum {
			maximum = total
		}
	}

	return maximum
}
