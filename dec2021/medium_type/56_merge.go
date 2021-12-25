package medium_type

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	var res [][]int

	if len(intervals) <= 1 {
		return intervals
	}

	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		lastEnd := res[len(res)-1][1]

		if intervals[i][0] <= lastEnd {
			res[len(res)-1][1] = max(lastEnd, intervals[i][1])
		} else {
			res = append(res, []int{intervals[i][0], intervals[i][1]})
		}
	}

	return res
}
