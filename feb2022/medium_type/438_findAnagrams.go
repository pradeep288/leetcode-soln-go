package medium_type

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	var expected, cur map[uint8]int
	var res []int

	var createHashMap func(inp string) map[uint8]int

	createHashMap = func(inp string) map[uint8]int {
		m := make(map[uint8]int)
		for i := 0; i < len(inp); i++ {
			m[inp[i]] += 1
		}
		return m
	}
	expected = createHashMap(p)
	cur = createHashMap(s[0:len(p)])

	var isAnagram func(map[uint8]int, map[uint8]int) bool

	isAnagram = func(src map[uint8]int, target map[uint8]int) bool {
		if len(src) != len(target) {
			return false
		}
		for k, _ := range src {
			if src[k] != target[k] {
				return false
			}
		}
		return true
	}

	if isAnagram(cur, expected) {
		res = append(res, 0)
	}

	for i := 1; i < len(s)-len(p)+1; i++ {
		if cur[s[i-1]] >= 2 {
			cur[s[i-1]] -= 1
		} else {
			delete(cur, s[i-1])
		}
		if _, ok := cur[s[i+len(p)-1]]; ok {
			cur[s[i+len(p)-1]] += 1
		} else {
			cur[s[i+len(p)-1]] = 1
		}
		if isAnagram(cur, expected) {
			res = append(res, i)
		}
	}

	return res
}
