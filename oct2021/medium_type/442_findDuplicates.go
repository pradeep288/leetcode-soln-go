package medium_type

// https://www.youtube.com/watch?v=Q3mH_1xTJNg&ab_channel=TimothyHChang

func findDuplicates(nums []int) []int {
	var output []int

	for _, val := range nums {
		v := abs(val)
		if nums[v-1] < 0 {
			output = append(output, v)
		} else {
			nums[v-1] *= -1
		}
	}

	return output
}

func abs(v int) int {
	if v > 0 {
		return v
	}
	return v * -1
}
