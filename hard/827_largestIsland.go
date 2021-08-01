package hard

func largestIsland(grid [][]int) int {
	max, islandID, row, col,hashMap := 0,2,len(grid),len(grid[0]),make(map[int]int)
	if grid == nil || len(grid) ==0 || len(grid[0]) == 0{
		return max
	}
	// Identify Unique Island groups
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			if grid[i][j]==1{
				size := getIslandSize(grid, i, j, islandID)
				hashMap[islandID]=size
				max = maxI(size, max)
				islandID += 1
			}
		}
	}
	directions := [][]int{{1,0},{0,1},{-1,0},{0,-1}}
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			var visited []int
			if grid[i][j] == 0{
				size := 1
				for _, direction := range directions{
					nextRow := i+ direction[0]
					nextCol := j+ direction[1]
					if nextRow <0|| nextCol <0 || nextRow >=row || nextCol >=col || islandVisited(grid[nextRow][nextCol], visited){
						continue
					} else {
						visited = append(visited, grid[nextRow][nextCol])
						size += hashMap[grid[nextRow][nextCol]]
					}
				}
				max = maxI(size,max)
			}
		}
	}
	return max
}

func getIslandSize(grid [][]int, r int, c int, id int)int{
	if r<0|| c<0 || r >=len(grid) || c >=len(grid[0]) || grid[r][c]!=1{
		return 0
	}
	grid[r][c] = id
	left := getIslandSize(grid, r, c-1, id)
	right := getIslandSize(grid, r, c+1, id)
	up := getIslandSize(grid, r-1, c, id)
	down := getIslandSize(grid, r+1, c, id)
	return left + right + up + down + 1
}

func maxI(x int, y int) int {
	if x>y{
		return x
	}
	return y
}

func islandVisited(id int, visited []int) bool{
	for _, v:=range visited{
		if v == id{
			return true
		}
	}
	return false
}