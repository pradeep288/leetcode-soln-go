package easy_type

import (
	"sort"
)

func findTheDifference(s string, t string) byte {
	if len(s) == 0 {
		return byte(t[0])
	}

	var sortString func(str string) string
	sortString = func(str string) string {
		r := []rune(str)
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		return string(r)
	}

	s = sortString(s)
	t = sortString(t)

	for i := 0; i < len(t)-1; i++ {
		if s[i] != t[i] {
			return byte(t[i])
		}
	}
	return byte(t[len(t)-1])
}
