package easy_type

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	var pToW map[rune]string
	var wToP map[string]rune

	pToW = make(map[rune]string)
	wToP = make(map[string]rune)

	count := 0
	for _, v := range pattern {
		if _, ok := pToW[v]; !ok {
			pToW[v] = words[count]
		} else if pToW[v] != words[count] {
			return false
		}

		if _, ok := wToP[words[count]]; !ok {
			wToP[words[count]] = v
		} else if wToP[words[count]] != v {
			return false
		}
		count++
	}
	return true
}
