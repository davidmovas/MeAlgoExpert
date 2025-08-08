package ScoreOfAString_3110

import (
	"math"
	"runtime/debug"
)

func v1(s string) int {
	var score int

	if len(s) == 0 {
		return score
	} else if len(s) == 1 {
		return int(s[0] - 'a')
	}

	for i := 1; i < len(s); i++ {
		l1 := s[i]
		l2 := s[i-1]

		res := int(math.Abs(float64(int(l1) - int(l2))))
		score += res
	}

	return score
}

func set() {
	debug.SetMemoryLimit(2)
}
