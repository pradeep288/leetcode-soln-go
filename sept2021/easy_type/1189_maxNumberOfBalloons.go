package easy_type

import (
	"math"
)

func maxNumberOfBalloons(text string) int {
	letterCnt := make(map[rune]int)

	//count all the letters in string
	for _, rn := range text {
		letterCnt[rn]++
	}

	// divide l and o since they occur twice in "balloon"
	letterCnt['l'] /= 2
	letterCnt['o'] /= 2

	// search for least seen character
	output := math.MaxInt32
	for _, rn := range "balon" {
		if letterCnt[rn] < output {
			output = letterCnt[rn]
		}
	}

	return output
}
