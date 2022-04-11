package medium_type

import "sort"

func deleteAndEarn(nums []int) int {

	getFeqCount := func(lst []int) map[int]int {
		freq := make(map[int]int)

		for _, v := range lst {
			freq[v] += 1
		}
		return freq
	}

	removeDuplicates := func(hashMap map[int]int) []int {
		var list []int
		for k, _ := range hashMap {
			list = append(list, k)
		}
		return list
	}

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	var dp []int
	count := getFeqCount(nums)
	tempList := removeDuplicates(count)

	if len(tempList) == 1 {
		return tempList[0] * count[tempList[0]]
	}
	sort.Ints(tempList)
	// Case 1: when the index is 0
	dp = append(dp, tempList[0]*count[tempList[0]])
	// Case 2: When the index is 1
	if tempList[1] == tempList[0]+1 {
		dp = append(dp, max(dp[0], tempList[1]*count[tempList[1]]))
	} else {
		dp = append(dp, dp[0]+tempList[1]*count[tempList[1]])
	}
	// Case 3: When the index is >1
	for i := 2; i < len(tempList); i++ {
		if tempList[i] == tempList[i-1]+1 {
			dp = append(dp, max(dp[i-1], dp[i-2]+tempList[i]*count[tempList[i]]))
		} else {
			dp = append(dp, dp[i-1]+tempList[i]*count[tempList[i]])
		}
	}
	return dp[len(dp)-1]
}
