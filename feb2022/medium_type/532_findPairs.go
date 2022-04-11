package medium_type

func findPairs(nums []int, k int) int {
	freqMap := make(map[int]int)
	for _, val := range nums {
		freqMap[val] += 1
	}

	var count int
	if k == 0 {
		for _, val := range freqMap {
			if val > 1 {
				count++
			}
		}
	} else {
		for key, _ := range freqMap {
			if _, ok := freqMap[key+k]; ok {
				count++
			}
		}
	}
	return count
}
