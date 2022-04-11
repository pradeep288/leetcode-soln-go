package easy_type

func isValid(s string) bool {
	hashMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	if len(s) == 0 {
		return true
	}
	var stack []string
	for _, char := range s {
		switch string(char) {
		case "(", "[", "{":
			stack = append(stack, string(char))
		default:
			if len(stack) > 0 && stack[len(stack)-1] == hashMap[string(char)] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, string(char))
			}
		}
	}
	return len(stack) == 0
}
