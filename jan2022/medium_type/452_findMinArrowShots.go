package medium_type

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	// Sort the Points based on the start Index
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	var min func(int, int) int

	min = func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	overlaps := 1

	prev := points[0]

	for _, point := range points {
		if prev[1] >= point[0] {
			prev[1] = min(prev[1], point[1])
		} else {
			overlaps++
			prev = point
		}
	}

	return overlaps
}
