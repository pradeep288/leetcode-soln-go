package medium_type

func rotate(nums []int, k int) {

	low, high := 0, len(nums)-1
	for low < high {
		nums[low], nums[high] = nums[high], nums[low]
		low += 1
		high -= 1
	}

	k = k % len(nums)

	low, high = 0, k-1
	for low < high {
		nums[low], nums[high] = nums[high], nums[low]
		low += 1
		high -= 1
	}

	low, high = k, len(nums)-1
	for low < high {
		nums[low], nums[high] = nums[high], nums[low]
		low += 1
		high -= 1
	}
}
