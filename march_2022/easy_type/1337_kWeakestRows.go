package easy_type

import (
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	freqMap := make(map[int][]int)
	for i, m := range mat {
		c := 0
		for _, v := range m {
			c += v
		}
		freqMap[c] = append(freqMap[c], i)
	}

	var soldiersCount []int
	for soldierCount, _ := range freqMap {
		soldiersCount = append(soldiersCount, soldierCount)
	}

	sort.Ints(soldiersCount)

	var res []int
	for _, soldierCount := range soldiersCount {
		sort.Ints(freqMap[soldierCount])
		for _, row := range freqMap[soldierCount] {
			res = append(res, row)
		}

	}
	return res[:k]
}
