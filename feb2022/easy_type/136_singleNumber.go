package easy_type

func singleNumber(nums []int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}
