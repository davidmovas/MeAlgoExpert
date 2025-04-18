package Candy_135

func v1(ratings []int) int {
	if len(ratings) == 1 {
		return 1
	}

	var candies = make([]int, len(ratings))

	for i := range candies {
		candies[i] += 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	var result int
	for _, c := range candies {
		result += c
	}

	return result
}
