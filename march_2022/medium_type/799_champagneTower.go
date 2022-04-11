package medium_type

func champagneTower(poured int, query_row int, query_glass int) float64 {
	var glasses []float64

	min := func(i, j float64) float64 {
		if i < j {
			return i
		}
		return j
	}

	glasses = append(glasses, float64(poured))

	for i := 0; i < query_row; i++ {
		temp := make([]float64, len(glasses)+1)
		for idx, v := range glasses {
			pour := (v - 1) / 2
			if pour > 0 {
				temp[idx] += pour
				temp[idx+1] += pour
			}
		}
		glasses = temp
	}

	return min(1, glasses[query_glass])
}
