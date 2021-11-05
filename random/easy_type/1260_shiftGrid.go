package easy_type

func shiftGrid(grid [][]int, k int) [][]int {
	var temp []int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			temp = append(temp, grid[i][j])
		}
	}

	start := len(temp) - (k % len(temp))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if start < len(temp) {
				grid[i][j] = temp[start]
				start++
			}
		}
	}
	var count, counter int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if count < (k % len(temp)) {
				count++
			} else {
				grid[i][j] = temp[counter]
				counter++
			}
		}
	}

	return grid
}
