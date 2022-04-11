package easy_type

/*func addDigits(num int) int {
	var digitalRoot int
	for num > 0 {
		digitalRoot += num % 10
		num /= 10

		if num == 0 && digitalRoot > 9 {
			num = digitalRoot
			digitalRoot = 0
		}
	}
	return digitalRoot
}*/

func addDigits(num int) int {
	if num < 9 {
		return num
	}
	switch num % 9 {
	case 0:
		return 9
	default:
		return num % 9
	}
}
