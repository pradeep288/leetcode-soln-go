package easy_type

func maxCount(m int, n int, ops [][]int) int {
	minR, minC := m, n

	for _, val:=range ops{
		minR = min(minR, val[0])
		minC = min(minC, val[1])
	}
	return minC * minR
}

func min(x, y int)int{
	if x<y{
		return x
	}
	return y
}