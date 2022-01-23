package medium_type

func sequentialDigits(low int, high int) []int {
	var q, res []int
	for i := 1; i <= 9; i++ {
		q = append(q, i)
	}

	for len(q) > 0 {
		curNum := q[0]
		q = q[1:]
		if curNum >= low && curNum <= high {
			res = append(res, curNum)
		}
		nextNum := curNum*10 + curNum%10 + 1
		if nextNum <= high && curNum%10 != 9 {
			q = append(q, nextNum)
		}
	}
	return res
}
