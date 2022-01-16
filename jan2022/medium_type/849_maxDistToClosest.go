package medium_type

/*
Algorithm:
Step1: Traverse From Left to Right and update the dpLtoR
Step2: Traverse From Right to Left and update the dpRtoL
Step3: Traverse both dpLtoR, dpRtoL and get the minimum(temp) at that index and update the res,
if temp is greater than temp
Step4: return res

Ex:
Input: [1,0,0,0,1,0,1]
Step1: [0 1 2 3 0 1 0]
Step2: [0 3 2 1 0 1 0]
Step3:      2
Step4: 2
*/
import (
	"math"
)

func maxDistToClosest(seats []int) int {
	dpLtoR := make([]int, len(seats))
	dpRtoL := make([]int, len(seats))

	// Update dpLtoR with results, when you traverse from left to right
	// Handle special case when seats[0] is 0
	if seats[0] == 0 {
		dpLtoR[0] = math.MaxInt16
	}
	for i := 1; i < len(seats); i++ {
		if seats[i] == 0 {
			dpLtoR[i] = dpLtoR[i-1] + 1
		}
	}

	// Update dpRtoL with results, when you traverse from left to right
	// Handle special case when seats[len(seats)-1] is 0
	if seats[len(seats)-1] == 0 {
		dpRtoL[len(seats)-1] = math.MaxInt16
	}
	for i := len(seats) - 2; i >= 0; i-- {
		if seats[i] == 0 {
			dpRtoL[i] = dpRtoL[i+1] + 1
		}
	}

	var res int
	res = math.MinInt16
	for i := 0; i < len(seats); i++ {
		if seats[i] != 1 {
			temp := min(dpLtoR[i], dpRtoL[i])
			if temp > res {
				res = temp
			}
		}
	}
	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
