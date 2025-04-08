package TandemBicycle

import algox "algoexpert/additional"

func TendemBicycle_V1(redRiders, blueRiders []int, fastest bool) int {
	var result int

	algox.Sort(redRiders, func(a, b int) bool { return a < b })
	algox.Sort(blueRiders, func(a, b int) bool { return a > b })

	for idx, _ := range redRiders {
		if redRiders[idx] > blueRiders[idx] {
			result += redRiders[idx]
		} else {
			result += blueRiders[idx]
		}
	}

	return result
}
