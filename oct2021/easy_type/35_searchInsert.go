package easy_type

import "fmt"

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	var mid int

	for low < high {
		mid = (low + high) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
		fmt.Println(low, high)
	}

	if target > nums[mid] {
		return mid + 1
	} else {
		return mid - 1
	}

}
