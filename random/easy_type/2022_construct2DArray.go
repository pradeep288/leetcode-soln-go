package easy_type

import "sort"

func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}

	res, idx := make([][]int, m), 0
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < m; j++ {
			res[i][j] = original[idx]
			idx++
		}
	}
	
	return res
}
