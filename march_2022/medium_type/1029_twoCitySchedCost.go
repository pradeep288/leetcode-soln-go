//  https://www.youtube.com/watch?v=d-B_gk_gJtQ
package medium_type

import "sort"

func twoCitySchedCost(costs [][]int) int {
	var diff [][]int
	for _, cost := range costs {
		diff = append(diff, []int{cost[1] - cost[0], cost[0], cost[1]})
	}
	sort.Slice(diff, func(i, j int) bool {
		return diff[i][0] < diff[j][0]
	})

	var res int
	for i := 0; i < len(diff); i++ {
		if i < len(costs)/2 {
			res += diff[i][2]
		} else {
			res += diff[i][1]
		}
	}
	return res
}
