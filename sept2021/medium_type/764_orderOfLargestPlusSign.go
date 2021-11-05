package medium_type

import "math"

func orderOfLargestPlusSign(n int, mines [][]int) int {
	var grid [][]int

	for i := 0; i < n; i++ {
		grid = append(grid, make([]int, n))
		for j := 0; j < n; j++ {
			grid[i][j] = 1
		}
	}

	for _, mine := range mines {
		grid[mine[0]][mine[1]] = 0
	}

	var dp [][]int
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, n))
		for j := 0; j < n; j++ {
			dp[i][j] = n
		}
	}

	var curSum int
	for i := 0; i < n; i++ {
		curSum = 0
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				curSum = 0
			} else {
				curSum += 1
			}
			dp[i][j] = min764(dp[i][j], curSum)
		}

		curSum = 0
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 0 {
				curSum = 0
			} else {
				curSum += 1
			}
			dp[i][j] = min764(dp[i][j], curSum)
		}
	}

	for i := 0; i < n; i++ {
		curSum = 0
		for j := 0; j < n; j++ {
			if grid[j][i] == 0 {
				curSum = 0
			} else {
				curSum += 1
			}
			dp[j][i] = min764(dp[j][i], curSum)
		}

		curSum = 0
		for j := n - 1; j >= 0; j-- {
			if grid[j][i] == 0 {
				curSum = 0
			} else {
				curSum += 1
			}
			dp[j][i] = min764(dp[j][i], curSum)
		}
	}

	var ans int
	ans = math.MinInt8
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans = max764(ans, dp[i][j])
		}
	}
	return ans
}

func min764(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max764(a, b int) int {
	if a > b {
		return a
	}
	return b
}
