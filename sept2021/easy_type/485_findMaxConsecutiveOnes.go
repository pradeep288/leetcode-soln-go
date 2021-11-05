package easy_type

func findMaxConsecutiveOnes(nums []int) int {
	var res, fp int

	for fp < len(nums) {
		var count int
		for fp < len(nums) {
			if nums[fp] == 1 {
				count += 1
				fp += 1
			} else {
				break
			}
		}
		res = max(count, res)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
