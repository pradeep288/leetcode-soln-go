package medium_type

func maxProduct(nums []int) int {
	var dp []int
	dp = append(dp, nums[0])
	var res int
	res = dp[0]
	for i := 1; i < len(nums); i++ {
		m := max(dp[i-1]*nums[i], nums[i])
		res = max(res, m)
		dp = append(dp, m)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
