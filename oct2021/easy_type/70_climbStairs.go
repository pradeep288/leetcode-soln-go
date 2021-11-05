package easy_type

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	a, b, count, res := 1, 2, n-2, 0

	for count > 0 {
		res = a + b
		a, b = b, res
		count--
	}

	return res

}
