package medium_type

func getSmallestString(n int, k int) string {
	res := make([]rune, n)
	for i := 0; i < n; i++ {
		res[i] = 'a'
	}

	k = k - n
	for i := n - 1; i >= 0; i-- {
		switch {
		case k > 25:
			res[i] = 'z'
			k = k - 25
		default:
			res[i] = 'a' + rune(k)
			return string(res)
		}
	}
	return ""
}
