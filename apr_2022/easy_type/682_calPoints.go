package easy_type

import "strconv"

func calPoints(ops []string) int {
	var stack []int
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		case "D":
			stack = append(stack, 2*stack[len(stack)-1])
		case "C":
			stack = stack[:len(stack)-1]
		default:
			val, _ := strconv.Atoi(ops[i])
			stack = append(stack, val)
		}
	}

	var res int
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}
