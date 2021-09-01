// https://www.youtube.com/watch?v=24DvdE-XP5E&ab_channel=TimothyHChang

package easy_type

func matrixReshape(mat [][]int, r int, c int) [][]int {
	res:=make([][]int,r)
	temp:=make([]int, r*c)


	count:=0
	rows, cols := len(mat), len(mat[0])

	if r*c != rows*cols{
		return mat
	}

	for i:=0;i<rows;i++{
		for j:=0;j<cols;j++{
			temp[count] =  mat[i][j]
			count+=1
		}
	}
	count = 0
	for i:=0;i<r;i++{
		res[i] = make([]int, c)
		for j:=0;j<c;j++{
			res[i][j] = temp[count]
			count += 1
		}
	}

	return res
}
