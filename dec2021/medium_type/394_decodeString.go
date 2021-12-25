package medium_type

import (
	"strings"
	"unicode"
)

func decodeString(s string) string {
	stackNums := make([]int, 0)
	stackStr := make([]string, 0)
	var res string
	var num int

	for i := 0; i < len(s); i++ {
		switch {
		case unicode.IsDigit(rune(s[i])):
			num = 10*num + int(s[i]) - '0'
		case s[i] == '[':
			stackNums = append(stackNums, num)
			num = 0
			stackStr = append(stackStr, res)
			res = ""
		case s[i] == ']':
			tmp := stackStr[len(stackStr)-1]
			stackStr = stackStr[:len(stackStr)-1]
			count := stackNums[len(stackNums)-1]
			stackNums = stackNums[:len(stackNums)-1]
			res = tmp + strings.Repeat(res, count)
		default:
			res += string(s[i])
		}

	}
	return res
}
