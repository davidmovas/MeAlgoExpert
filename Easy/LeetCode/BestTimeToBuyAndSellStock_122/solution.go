package BestTimeToBuyAndSellStock1_121

import "math"

// prices = [7,1,5,3,6,4]

// TODO: NOT MY SOLUTION
func v1(prices []int) int {
	curHold, curNotHold := math.MinInt64, 0

	for _, price := range prices {

		prevHold, prevNotHold := curHold, curNotHold

		// either keep hold, or buy in stock today at stock price
		curHold = Max(prevHold, prevNotHold-price)

		// either keep not-hold, or sell out stock today at stock price
		curNotHold = Max(prevNotHold, prevHold+price)
	}

	return curNotHold
}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
