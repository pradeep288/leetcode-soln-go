package easy_type

func largestAltitude(gain []int) int {
	maxAltitude := 0
	totalGain := 0

	for _, g := range gain {
		totalGain += g
		if totalGain > maxAltitude {
			maxAltitude = totalGain
		}
	}

	return maxAltitude
}
