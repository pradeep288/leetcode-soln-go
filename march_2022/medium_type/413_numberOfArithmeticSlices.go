package medium_type

func numberOfArithmeticSlices(nums []int) int {
	dp := make([]int, len(nums))
	var total int
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
			total += dp[i]
		}
	}
	return total
}
