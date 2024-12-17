package SmallestDifference

import (
	"math"
	"sort"
)

func SmallestDifference(first, second []int) []int {
	firstIdx, secondIdx := 0, 0
	currentSmallestSum := math.MaxInt

	sort.Ints(first)
	sort.Ints(second)

	for firstIdx < len(first) && secondIdx < len(second) {
		tempFirst := first[firstIdx]
		tempSecond := second[secondIdx]

		if tempFirst < 0 {
			tempFirst *= -1
		}

		if tempSecond < 0 {
			tempSecond *= -1
		}

		var difference int

		if tempFirst > tempSecond {
			difference = tempSecond - tempFirst
		} else {
			difference = tempFirst - tempSecond
		}

		if difference < currentSmallestSum {
			currentSmallestSum = difference
		}
	}

	return []int{first[firstIdx], second[secondIdx]}
}
