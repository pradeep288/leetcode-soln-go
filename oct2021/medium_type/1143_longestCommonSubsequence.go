package medium_type

func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == text2 {
		return len(text1)
	}

	dp := make([][]int, len(text1)+1)

	for j := 0; j < len(text1)+1; j++ {
		dp[j] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
  "" a b c d e
0  0 0 0 0 0 0
a  0 1 1 1 1 1
c  0 1 1 2 2 2
e  0 1 1 2 2 3
*/
