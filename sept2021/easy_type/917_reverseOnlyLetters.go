package easy_type

import "unicode"

func reverseOnlyLetters(s string) string {
	chars := []rune(s)
	var start, end int

	start = 0
	end = len(s) - 1

	for start < end {
		if !isEnglishLetter(rune(chars[start])) {
			start++
			continue
		}
		if !isEnglishLetter(rune(chars[end])) {
			end--
			continue
		}

		chars[start], chars[end] = chars[end], chars[start]
		end--
		start++
	}

	return string(chars)
}

func isEnglishLetter(val rune) bool {
	if unicode.IsLetter(val) {
		return true
	}
	return false
}
