package easy_type

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	hashMap := make(map[string]int)
	hashMap["type"] = 0
	hashMap["color"] = 1
	hashMap["name"] = 2

	var res int
	for _, v := range items {
		if v[hashMap[ruleKey]] == ruleValue {
			res++
		}
	}
	return res
}
