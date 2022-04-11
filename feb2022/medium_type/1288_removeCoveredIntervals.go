package medium_type

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	unConveredIntervals := len(intervals)
	for i := 0; i < len(intervals)-1; i++ {
		curNode := intervals[i]
		for j := 0; j < len(intervals); j++ {
			if j != i {
				if intervals[j][0] <= curNode[0] && intervals[j][i] >= curNode[1] {
					unConveredIntervals -= 1
					break
				}
			}
		}
	}

	return unConveredIntervals
}
