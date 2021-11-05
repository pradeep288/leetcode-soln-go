// https://www.youtube.com/watch?v=aZuQXkO0-XA&t=4s&ab_channel=Pepcoding
package medium_type

import "math"

func numSquares(n int) int {
	var dp []int
	dp = append(dp, 0, 1, 2, 3)

	for i := 4; i <= n; i++ {
		minVal := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			rem := i - j*j
			if dp[rem] < minVal {
				minVal = dp[rem]
			}
		}
		dp = append(dp, minVal+1)
	}
	return dp[n]
}
