package easy_type

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	hashMap := make(map[int]int)

	for i, v := range score {
		hashMap[v] = i
	}

	sort.Ints(score)

	res := make([]string, len(hashMap))

	for i := len(score) - 1; i >= 0; i-- {
		switch len(score) - i {
		case 1:
			res[hashMap[score[i]]] = "Gold Medal"
		case 2:
			res[hashMap[score[i]]] = "Silver Medal"
		case 3:
			res[hashMap[score[i]]] = "Bronze Medal"
		default:
			res[hashMap[score[i]]] = strconv.Itoa(len(score) - i)
		}
	}

	return res
}
