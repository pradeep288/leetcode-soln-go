package easy_type

func shiftGrid(grid [][]int, k int) [][]int {
	var tmp []int

	// Convert into 1d array
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			tmp = append(tmp, grid[i][j])
		}
	}

	k = k % len(tmp)

	tmp = append(tmp[len(tmp)-k:], tmp[:len(tmp)-k]...)
	var idx int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = tmp[idx]
			idx++
		}
	}
	return grid
}
