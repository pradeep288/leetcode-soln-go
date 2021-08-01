package medium_type

func updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 {
		return mat
	}

	row,col:=len(mat),len(mat[0])
	var visited [][]bool
	var q [][]int

	for i:=0;i<row;i++{
		var row []bool
		for j:=0;j<col;j++{
			if mat[i][j]==0{
				q = append(q,[]int{i,j})
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		visited = append(visited,row)
	}
	dirs := [][]int{{1,0},{0,1},{-1,0},{0,-1}}

	for len(q)>0{
		size := len(q)
		for i:=0;i<size;i++{
		curCell := q[0]
		curX,curY:=curCell[0],curCell[1]
		for _,dir:=range dirs{
			nextX:=curX+dir[0]
			nextY:=curY+dir[1]
			if nextX<0 || nextX > row-1 || nextY <0 || nextY>col-1 || visited[nextX][nextY]{
				continue
			}
			mat[nextX][nextY] = mat[curX][curY] + 1
			visited[nextX][nextY] =true
		}
		}
		q = q[1:]
	}

	return mat
}