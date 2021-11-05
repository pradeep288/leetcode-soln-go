package easy_type

import "math"

func titleToNumber(columnTitle string) int {
	var res float64

	for i := 0; i < len(columnTitle); i++ {
		p := len(columnTitle) - 1 - i
		v := math.Pow(26, float64(p))
		res += v * float64(rune(columnTitle[i])-'A'+1)
	}

	return int(res)
}
