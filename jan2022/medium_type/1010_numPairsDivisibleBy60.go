package medium_type

func numPairsDivisibleBy60(time []int) int {
	var res int
	freq := make(map[int]int)

	// Obtain time[i]%60 and update the hashmap with time[i]%60 as key and its count as val
	for i := 0; i < len(time); i++ {
		mod := time[i] % 60
		freq[mod]++
	}

	for k, v := range freq {
		switch {
		// if k==0/30 the number of pairs is n*n-1/2
		case k == 0 || k == 30:
			res += (freq[k] * (freq[k] - 1)) / 2
		case k <= 29:
			if _, ok := freq[60-k]; ok {
				res += v * freq[60-k]
			}
		}
	}

	return res
}
