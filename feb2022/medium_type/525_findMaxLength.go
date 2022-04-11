package medium_type

func findMaxLength(nums []int) int {
	var sum, maxLen int
	var hashMap map[int]int
	hashMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum -= 1
		} else {
			sum += 1
		}
		if sum == 0 {
			maxLen = i + 1
		} else {
			if _, ok := hashMap[sum]; ok {
				if i-hashMap[sum] > maxLen {
					maxLen = i - hashMap[sum]
				}
			} else {
				hashMap[sum] = i
			}
		}

	}
	return maxLen
}
