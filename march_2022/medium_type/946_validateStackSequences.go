package medium_type

import "fmt"

// pushed = [1,2,3,4,5]
// popped = [4,5,3,2,1]
func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	ctrPush := 0
	ctrPop := 0
	for true {
		if ctrPush < len(pushed) && pushed[ctrPush] != popped[ctrPop] {
			stack = append(stack, pushed[ctrPush])
			ctrPush += 1
		} else if stack[len(stack)-1] == popped[ctrPop] {
			stack = stack[:len(stack)-1]
			ctrPush += 1
			ctrPop += 1
		} else {
			break
		}
	}
	fmt.Println(ctrPush, len(pushed), ctrPop, len(popped), len(stack))

	return ctrPush == len(pushed) && ctrPop == len(popped) && len(stack) == 0
}
