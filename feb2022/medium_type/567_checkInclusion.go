package medium_type

/*
import (
	"reflect"
	"sort"
)

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	src := []uint8(s1)
	sort.Slice(src, func(i, j int) bool {
		return src[i] < src[j]
	})

	for i := 0; i <= len(s2)-len(s1); i++ {
		tmp := []uint8(s2[i : i+len(s1)])
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		if reflect.DeepEqual(src, tmp) {
			return true
		}
	}

	return false
}
*/

import "reflect"

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	var createHashMap func(string) map[int32]int
	createHashMap = func(s string) map[int32]int {
		hashMap := make(map[int32]int)
		for _, val := range s {
			hashMap[val] += 1
		}
		return hashMap
	}

	src := createHashMap(s1)

	for i := 0; i <= len(s2)-len(s1); i++ {
		tmp := createHashMap(s2[i : i+len(s1)])
		if reflect.DeepEqual(tmp, src) {
			return true
		}
	}
	return false
}
