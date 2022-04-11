package medium_type

func subarraySum(nums []int, k int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		curSum := 0
		for j := i; j < len(nums); j++ {
			if curSum+nums[j] == k {
				count++
			}
			curSum += nums[j]
		}
	}
	return count
}
