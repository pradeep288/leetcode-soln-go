package medium_type

import "sort"

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freqMap[nums[i]] += 1
	}

	var keys []int
	for key, _ := range freqMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return freqMap[keys[i]] > freqMap[keys[j]]
	})

	return keys[:k]
}
