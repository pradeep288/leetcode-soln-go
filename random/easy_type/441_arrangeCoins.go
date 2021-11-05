package easy_type

func arrangeCoins(n int) int {
	var usedCoins, completeRows, count int

	for true {
		count += 1
		if n-usedCoins < count {
			break
		}
		usedCoins += count
		completeRows++
	}
	return completeRows
}
