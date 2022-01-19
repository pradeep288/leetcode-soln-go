package hard

import "math"

var elements map[rune]int

func minWindow(s string, t string) string {
	elements = make(map[rune]int, 0)
	var minWinSubStr string
	minLen := math.MaxInt16

	if len(s) < len(t) {
		return minWinSubStr
	}

	for _, val := range []rune(t) {
		if _, ok := elements[val]; ok {
			elements[val] += 1
		} else {
			elements[val] = 1
		}
	}

	sp := 0
	fp := len(t) - 1

	for sp <= fp && fp <= len(s) {
		if checkTargetFound(s[sp:fp]) {
			for true {
				if len(s[sp:fp]) < minLen {
					minLen = len(s[sp:fp])
					minWinSubStr = s[sp:fp]
				}
				sp += 1
				if !checkTargetFound(s[sp:fp]) {
					break
				}
			}
		} else {
			fp += 1
		}
	}

	return minWinSubStr
}

func checkTargetFound(s string) bool {
	temp := make(map[rune]int, 0)
	for k, v := range elements {
		temp[k] = v
	}

	for _, v := range []rune(s) {
		if _, ok := temp[v]; ok {
			temp[v] -= 1
			if temp[v] == 0 {
				delete(temp, v)
			}
		}
	}

	return len(temp) == 0
}
