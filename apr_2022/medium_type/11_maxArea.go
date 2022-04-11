package medium_type

func maxArea(height []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	l, r, maxArea := 0, len(height)-1, 0

	for l < r {
		curArea := (r - l) * min(height[l], height[r])
		maxArea = max(maxArea, curArea)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}
