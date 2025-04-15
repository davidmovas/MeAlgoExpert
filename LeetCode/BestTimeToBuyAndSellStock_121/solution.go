package BestTimeToBuyAndSellStock_121

func v1(prices []int) int {
	var result = 0

	for i := 0; i <= len(prices)-1; i++ {
		for j := i + 1; j <= len(prices)-1; j++ {
			profit := prices[j] - prices[i]
			if profit > result {
				result = profit
			}
		}
	}

	return result
}

func v2(prices []int) int {
	result, minimal, maximal := 0, prices[0], prices[0]

	for _, price := range prices {
		if price < minimal {
			minimal = price
			maximal = price
		} else if price > maximal {
			maximal = price
		}

		if (maximal - minimal) > result {
			result = maximal - minimal
		}
	}

	return result
}
