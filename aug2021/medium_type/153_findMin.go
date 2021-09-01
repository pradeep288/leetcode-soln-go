package medium_type

func findMin(nums []int) int {
	var l, h, m int
	l = 0
	h = len(nums) - 1
	m = (l + h) / 2

	for l < h {

		if nums[m] < nums[h] {
			h = m
		} else {
			l = m + 1
		}
		m = (l + h) / 2
	}
	return nums[l]
}
