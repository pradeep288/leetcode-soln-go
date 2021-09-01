package medium_type

// https://www.youtube.com/watch?v=WhKrPDS1VAg&ab_channel=code_reportcode_report
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	start, end, res  := -1, -1, 0

	for i:=0;i<len(nums);i++{
		// Handle Reset Value
		if nums[i] > right {
			start = i
		}

		// Handle reset value and needed value
		if nums[i] >= left{
			end = i
		}


		res += end - start
	}

	return res

}