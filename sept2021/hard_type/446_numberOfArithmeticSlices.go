package hard_type

import (
	"math"
)

// 1. DP Approach
// References: https://www.youtube.com/watch?v=TMN268YMWzg&ab_channel=CodingDecoded
func numberOfArithmeticSlices(nums []int) int {
	var ans int
	var dpMap []map[int]int

	// for every element nums[i] create a hashmap,
	// which stores diff of nums[j]- nums[i] where j<i
	for i := 0; i < len(nums); i++ {
		dpMap = append(dpMap, make(map[int]int))
		var n1, n2 int
		for j := 0; j < i; j++ {
			diff := nums[j] - nums[i]
			if diff <= math.MinInt32 || diff >= math.MaxInt32 {
				continue
			}

			// if the diff already exists on the nums[j] get the count of diff else set it 0.
			// this will be added to result
			if _, ok := dpMap[j][diff]; ok {
				n1 = dpMap[j][diff]
			} else {
				n1 = 0
			}

			// if the diff already exists on the nums[i] get the count of diff else set it 0.
			// this will be needed for updating the result.
			if _, ok := dpMap[i][diff]; ok {
				n2 = dpMap[i][diff]
			} else {
				n2 = 0
			}
			ans += n1
			dpMap[i][diff] = n1 + n2 + 1
		}
	}

	return ans
}
