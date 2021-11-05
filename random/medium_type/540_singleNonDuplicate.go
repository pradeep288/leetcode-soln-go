// https://www.youtube.com/watch?v=nMGL2vlyJk0&ab_channel=TECHDOSE
package medium_type

func singleNonDuplicate(nums []int) int {
	high := len(nums) - 1
	low := 0

	//check boundary condition
	if high == 0 {
		return nums[0]
	}
	if nums[high] != nums[high-1] {
		return nums[high]
	}
	if nums[low] != nums[low+1] {
		return nums[low]
	}

	// Run Binary Search
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
			return nums[mid]
		}

		if ((mid%2) == 0 && nums[mid] == nums[mid+1]) || ((mid%2) == 1 && nums[mid] == nums[mid-1]) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
