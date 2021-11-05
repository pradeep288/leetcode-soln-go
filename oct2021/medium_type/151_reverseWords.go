package medium_type

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)
	var res string

	for i := len(words) - 1; i >= 0; i-- {
		res = res + words[i] + " "
	}

	return strings.TrimSpace(res)
}
