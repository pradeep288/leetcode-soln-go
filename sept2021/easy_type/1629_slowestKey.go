package easy_type

import (
	"math"
	"sort"
)

type keyTest struct {
	key         rune
	releaseTime int
	duration    int
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	var tests []keyTest
	maxVal := math.MinInt16
	for i, v := range keysPressed {
		var duration int
		if i == 0 {
			duration = releaseTimes[i] - 0
		} else {
			duration = releaseTimes[i] - releaseTimes[i-1]
		}
		if duration > maxVal {
			maxVal = duration
		}

		tests = append(tests, keyTest{
			key:         v,
			releaseTime: releaseTimes[i],
			duration:    duration,
		})
	}

	sort.Slice(tests, func(i, j int) bool {
		if tests[i].duration == tests[j].duration {
			if tests[i].key < tests[j].key {
				return false
			}
			return true
		}
		if tests[i].duration == tests[j].duration {
			return true
		}
		return false
	})

	return byte(tests[0].key)
}
