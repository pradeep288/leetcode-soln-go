package easy_type

func canConstruct(ransomNote string, magazine string) bool {
	needed := make(map[rune]int)

	for _, v := range ransomNote {
		needed[v] += 1
	}

	for _, v := range magazine {
		if _, ok := needed[v]; ok {
			needed[v]--
			if needed[v] == 0 {
				delete(needed, v)
			}
		}
	}

	return len(needed) == 0
}
