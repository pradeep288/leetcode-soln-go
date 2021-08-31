package medium_type

func setZeroes(matrix [][]int)  {
	var cols, rows []int
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[0]);j++{
			if matrix[i][j] == 0{
				if ! valExists(i, rows){
					rows = append(rows, i)
				}
				if ! valExists(j, cols){
					cols = append(cols, j)
				}
			}
		}
	}

	for _,val:=range cols{
		for j:=0;j<len(matrix);j++{
			matrix[j][val] = 0
		}
	}

	for _,val:=range rows{
		for j:=0;j<len(matrix[0]);j++{
			matrix[val][j] = 0
		}
	}


}

func valExists(target int, s []int) bool{
	for i:=0;i<len(s);i++{
		if s[i] == target{
			return true
		}
	}

	return false
}
