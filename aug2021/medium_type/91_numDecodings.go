package medium_type

func numDecodings(s string) int {
	var dp []int
	dp = append(dp, 0)
	if s[0] == 48{
		dp = append(dp, 0)
	} else {
		dp = append(dp, 1)
	}
	for i:=2;i<len(s);i++{
		one := s[i]
		if one >48{
			dp[i] +=dp[i-1]
		}

	}

	return 0
}
