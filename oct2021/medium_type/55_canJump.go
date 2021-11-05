package medium_type

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	max_reach := nums[0]
	index := 1

	for max_reach >= index {
		if max_reach >= len(nums)-1 {
			return true
		}
		max_reach = max_55(max_reach, nums[index]+index)
		index++
	}
	return false
}

func max_55(a, b int) int {
	if a > b {
		return a
	}
	return b
}
