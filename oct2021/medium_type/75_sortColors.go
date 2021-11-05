// https://www.youtube.com/watch?v=uvB-Ns_TVis&ab_channel=NickWhite
package medium_type

func sortColors(nums []int) {
	nextZero, nextTwo, idx := 0, len(nums)-1, 0

	for idx <= nextTwo {
		if nums[idx] == 0 {
			nums[nextZero], nums[idx] = nums[idx], nums[nextZero]
			nextZero++
			idx++
		} else if nums[idx] == 1 {
			idx++
		} else {
			nums[nextTwo], nums[idx] = nums[idx], nums[nextTwo]
			nextTwo--
		}
	}
}
