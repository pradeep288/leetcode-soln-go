package easy_type

func maximumDifference(nums []int) int {
	max_diff := -1
	min_element := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min_element {
			min_element = nums[i]
		} else if max_diff < nums[i]-min_element {
			max_diff = nums[i] - min_element
		}
	}

	return max_diff
}
