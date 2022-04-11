package easy_type

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	var output []string
	var i, j int
	for i = 0; i < len(nums); {
		prev := nums[i]
		for j = i + 1; j < len(nums); j++ {
			if prev+1 != nums[j] {
				break
			}
			prev = nums[j]
		}
		if nums[i] == prev {
			output = append(output, strconv.Itoa(nums[i]))
			i = i + 1
		} else {
			output = append(output, strconv.Itoa(nums[i])+"->"+strconv.Itoa(prev))
			i = j
		}
	}

	return output
}
