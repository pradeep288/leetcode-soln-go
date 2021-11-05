package easy_type

func subtractProductAndSum(n int) int {
	var digits []int
	for n >= 10 {
		rem := n % 10
		digits = append(digits, rem)
		n = n / 10
	}
	digits = append(digits, n)

	sum := 0
	mul := 1

	for _, v := range digits {
		sum += v
		mul *= v
	}

	return mul - sum
}
