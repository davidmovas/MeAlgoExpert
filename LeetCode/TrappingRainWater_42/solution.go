package TrappingRainWater_42

func v1(height []int) int {
	var lMax, rMax, water = 0, 0, 0
	var left, right = 0, len(height) - 1

	for left < right {
		rMax = max(height[right], rMax)
		lMax = max(height[left], lMax)
		minH := min(lMax, rMax)

		if minH > height[left] {
			water += minH - height[left]
		}
		if minH > height[right] {
			water += minH - height[right]
		}

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return water
}
