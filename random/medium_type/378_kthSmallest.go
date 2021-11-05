package medium_type

import "sort"

func kthSmallest(matrix [][]int, k int) int {
	var nums []int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			nums = append(nums, matrix[i][j])
		}
	}
	sort.Ints(nums)
	return nums[k-1]
}
