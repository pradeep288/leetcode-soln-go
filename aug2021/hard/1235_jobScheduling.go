package hard

import (
	"math"
	"sort"
)

type job struct {
	startTime int
	endTime   int
	profit    int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	var jobs []job
	dp := []int{}

	for i := 0; i < len(startTime); i++ {
		jobs = append(jobs, job{
			startTime[i],
			endTime[i],
			profit[i],
		})
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].startTime < jobs[j].startTime
	})

	res := math.MinInt16
	for i := 0; i < len(jobs); i++ {
		dp = append(dp, jobs[i].profit)

	}

	for i := 1; i < len(jobs); i++ {
		j := 0
		for j < i {
			if jobs[i].startTime >= jobs[j].endTime {
				dp[i] = max1235(dp[i], dp[j]+jobs[i].profit)
			}
			j += 1
		}
	}

	for _, val := range dp {
		if res < val {
			res = val
		}
	}

	return res
}

func max1235(x, y int) int {
	if x > y {
		return x
	}
	return y
}
