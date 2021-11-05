// https://www.youtube.com/watch?v=OEW50g03mT0&ab_channel=NickWhite
package easy_type

func addBinary(a string, b string) string {
	var sum, carry int

	inp1 := []rune(a)
	inp2 := []rune(b)
	var out []int
	i, j := len(inp1)-1, len(inp2)-1
	for i >= 0 || j >= 0 {
		sum += carry
		if i >= 0 {
			sum += int(inp1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(inp2[j] - '0')
			j--
		}
		carry = sum / 2
		sum %= 2

		out = append(out, sum)
	}

	var res string
	if carry != 0 {
		out = append(out, carry)
	}

	for i := len(out) - 1; i >= 0; i-- {
		res += string(out[i] + '0')
	}

	return res
}
