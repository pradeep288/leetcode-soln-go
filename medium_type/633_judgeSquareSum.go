package medium_type

import "math"

func judgeSquareSum(c int) bool {

	var sq []int
	s := math.Sqrt(float64(c))

	for i:=0;i<=int(s)+1;i++{
		sq = append(sq, i*i)
	}

	low:=0
	high:=len(sq)-1

	for low<high{
		if sq[low]+ sq[high] == c{
			return true
		}

		if sq[low]+ sq[high] < c{
			low += 1
		} else{
			high -= 1
		}


	}
	return false
}
