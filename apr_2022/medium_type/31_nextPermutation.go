package medium_type

import "sort"

func nextPermutation(nums []int) {
	// Identify the index of last peak.
	lastPeakIdx := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			lastPeakIdx = i
			break
		}
	}

	// Case a: If the given input is in descending order.
	if lastPeakIdx == -1 {
		sort.Ints(nums)
		return
	}

	// case b: If the given input is not in descending order.
	idx := lastPeakIdx
	for i := lastPeakIdx; i < len(nums); i++ {
		if nums[i] > nums[lastPeakIdx-1] && nums[i] < nums[idx] {
			// Special Case: There exists an element between lastPeakIdx  and len(nums)-1
			// which is greater than the nums[lastPeakIdx-1] and less than the nums[idx]
			idx = i
		}
	}

	nums[idx], nums[lastPeakIdx-1] = nums[lastPeakIdx-1], nums[idx]
	sort.Ints(nums[lastPeakIdx:])
}
