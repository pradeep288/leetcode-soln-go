package easy_type

func islandPerimeter(grid [][]int) int {
	var perimeter int

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	for i := 0; i <= len(grid)-1; i++ {
		for j := 0; j <= len(grid[0])-1; j++ {
			if grid[i][j] == 1 {
				perimeter += 4
				if i > 0 && grid[i-1][j] == 1 {
					perimeter -= 2
				}

				if j > 0 && grid[i][j-1] == 1 {
					perimeter -= 2
				}
			}
		}
	}
	return perimeter
}
