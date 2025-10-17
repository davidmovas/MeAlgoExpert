package HappyNumber_202

func v1(n int) bool {
	if n <= 0 {
		return false
	}

	seen := make(map[int]bool)
	for n != 1 {
		if seen[n] {
			return false
		}
		seen[n] = true
		n = sumSquares(n)
	}
	return true
}

func sumSquares(x int) int {
	sum := 0
	for x > 0 {
		d := x % 10
		sum += d * d
		x /= 10
	}
	return sum
}
