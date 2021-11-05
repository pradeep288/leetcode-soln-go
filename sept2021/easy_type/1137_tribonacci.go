package easy_type

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	var cur int
	minus1 := 1
	minus2 := 1
	minus3 := 0

	for n > 2 {
		cur = minus1 + minus2 + minus3
		minus1, minus2, minus3 = cur, minus1, minus2
		n -= 1
	}

	return cur
}
