package hard

import "math"

func countOrders(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 6
	}
	dp := make([]int, n+1)
	dp = append(dp, 0, 1, 6)
	mod := math.Pow(10, 9) + 7

	for i := 3; i < n+1; i++ {
		oddNumber := 2*i - 1
		permutation := oddNumber * (oddNumber + 1) / 2
		dp[i] = permutation % int(mod)
	}
	return dp[n]
}
