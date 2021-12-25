package easy_type

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	var res [][]int
	var min int

	// Sort the Input array
	sort.Ints(arr)

	min = math.MaxInt16

	// find the  abs min diff
	for i := 0; i <= len(arr)-2; i++ {
		absDiff := abs(arr[i] - arr[i+1])
		if absDiff < min {
			min = absDiff
		}
	}

	//Find the pairs whose abs min diff is equal to min
	for i := 0; i <= len(arr)-2; i++ {
		absDiff := abs(arr[i] - arr[i+1])
		if absDiff == min {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}

	return res
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return x * -1
}
