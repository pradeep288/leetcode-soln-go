package hard

import "math"

func minCut(s string) int {
	size:=len(s)

	var dpArray [][]bool

	//Trivial case: single letter palindromes
	for i:=0;i<size;i++{
		var row []bool
		for j:=0;j<size;j++{
			if i==j{
				row = append(row, true)
			}else{
				row = append(row, false)
			}
		}
		dpArray = append(dpArray, row)
	}

	//Finding palindromes of two characters.
	for i:=0;i<size-1;i++{
		if s[i] == s[i+1]{
			dpArray[i][i+1] =true
		}
	}

	//Finding palindromes of length 3 to n
	for curLen:=3;curLen<=size;curLen++{
		for i:=0;i<size-curLen+1;i++{
			j:=i+curLen-1
			if s[i] == s[j] && dpArray[i+1][j-1]{
				dpArray[i][j] = true
			} else{
				dpArray[i][j] = false
			}
		}
	}

	cuts:=make([]int,size)
	for i:=0;i<size;i++ {
		temp:=math.MaxInt32
		if dpArray[0][i]{
			cuts[i] = 0
		} else {
			for j:=0;j<i;j++{
				if dpArray[j+1][i] && temp >cuts[j]+1{
					temp = cuts[j] + 1
				}
			}
			cuts[i] = temp
		}
	}

	return cuts[size-1]
}

