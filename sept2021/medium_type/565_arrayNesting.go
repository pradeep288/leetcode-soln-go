package medium_type

func arrayNesting(nums []int) int {
	maxSize := 0

	for i := 0; i < len(nums); i++ {
		size, next := 0, i
		for true {
			if nums[next] == -1 {
				break
			}
			a := nums[next]
			nums[next] = -1
			next = a
			size += 1
		}
		maxSize = max(maxSize, size)
	}

	return maxSize
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
