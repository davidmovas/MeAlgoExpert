package HIndex_274

import "sort"

func v1(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	var hIdx = 0

	for i, c := range citations {
		if c >= i+1 {
			hIdx = i + 1
		} else {
			break
		}
	}

	return hIdx
}
