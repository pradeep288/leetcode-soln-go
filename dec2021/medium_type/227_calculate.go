package medium_type

import "unicode"

func calculate(s string) int {
	var stack []int
	var curNum int
	var prevOp rune
	prevOp = '+'

	for i, v := range s {
		if unicode.IsDigit(v) {
			curNum = curNum*10 + int(v-'0')
		}
		if v == '+' || v == '-' || v == '*' || v == '/' || i == len(s)-1 {
			switch prevOp {
			case '+':
				stack = append(stack, curNum)
			case '-':
				stack = append(stack, -curNum)
			case '*':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, top*curNum)

			case '/':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, top/curNum)
			}
			curNum = 0
			prevOp = v
		}

	}

	var sum int
	for _, v := range stack {
		sum += v
	}

	return sum
}
