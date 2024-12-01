package MinimumWaitingTime

import "slices"

func MinimumWaitingTime_V1(queries []int) int {
	totalTime := 0

	slices.Sort(queries)

	for idx, duration := range queries {
		leftPart := len(queries) - (idx + 1)
		totalTime += duration * leftPart
	}

	return totalTime
}
