package hard_type

func winnerSquareGame(n int) bool {
	if n == 1 {
		return true
	}
	dp := make([]bool, n+1)

	dp[0] = false
	dp[1] = true
	dp[2] = false

	for i := 3; i < len(dp); i++ {
		j := 1
		for j*j <= i {
			if !dp[i-j*j] {
				dp[i] = true
				break
			}
			j += 1
		}
	}

	return dp[n]
}
