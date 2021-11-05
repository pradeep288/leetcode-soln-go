package easy_type

func arrangeCoins(n int) int {
	completeRows, available, count := 0, n, 0

	for true {
		count++
		if !(available-count >= 0) {
			break
		}
		available -= count
		completeRows++
	}

	return completeRows
}
