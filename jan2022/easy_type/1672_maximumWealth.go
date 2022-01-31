package easy_type

func maximumWealth(accounts [][]int) int {
	maxWealth := 0

	for i := 0; i < len(accounts); i++ {
		curWealth := 0
		for j := 0; j < len(accounts[i]); j++ {
			curWealth += accounts[i][j]
		}
		if curWealth > maxWealth {
			maxWealth = curWealth
		}
	}

	return maxWealth
}
