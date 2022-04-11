package medium_type

import "math"

func minDominoRotations(tops []int, bottoms []int) int {
	getMinRotations := func(arr1, arr2 []int) int {
		hashMap := make(map[int]int)
		minVal := math.MaxInt32
		for i := 0; i < len(arr1); i++ {
			if _, ok := hashMap[arr1[i]]; ok {
				continue
			}
			val, count, swapCount := arr1[i], 0, 0
			for j := 0; j < len(arr1); j++ {
				if i == j || arr1[j] == val {
					count++
				} else if arr2[j] == val {
					count++
					swapCount++
				}
			}
			if count == len(arr1) {
				if swapCount < minVal {
					minVal = swapCount
				}
			}
			hashMap[arr1[i]] = minVal
		}

		return minVal
	}

	topMin := getMinRotations(tops, bottoms)
	bottomMin := getMinRotations(bottoms, tops)

	if topMin == math.MaxInt32 && bottomMin == math.MaxInt32 {
		return -1
	} else {
		if topMin != math.MaxInt32 && bottomMin != math.MaxInt32 {
			if topMin < bottomMin {
				return topMin
			}
			return bottomMin
		} else if topMin == math.MaxInt32 {
			return bottomMin
		} else {
			return topMin
		}
	}
}
