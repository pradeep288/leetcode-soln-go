package medium_type

type info struct {
	idx int
	c   string
}

func minRemoveToMakeValid(s string) string {
	var stack []info
	for i, char := range s {
		switch string(char) {
		case "(":
			stack = append(stack, info{
				idx: i,
				c:   "(",
			})
		case ")":
			if len(stack) > 0 {
				if stack[len(stack)-1].c == "(" {
					stack = stack[:len(stack)-1]
					continue
				}
			}
			stack = append(stack, info{
				idx: i,
				c:   ")",
			})
		default:
			continue
		}
	}

	var res string
	var sp int

	for i, char := range s {
		if sp < len(stack) && i == stack[sp].idx {
			sp += 1
		} else {
			res += string(char)
		}
	}

	return res
}
