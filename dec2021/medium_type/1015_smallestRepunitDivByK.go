package medium_type

func smallestRepunitDivByK(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}
	if K == 1 {
		return 1
	}
	size := 1
	for v := 1; v != 0; size++ {
		if v = (v*10 + 1) % K; v == 0 {
			break
		}
	}
	return size + 1
}
