package easy_type

import "math"

func maxProfit(prices []int) int {
	minVal := math.MaxInt
	maxVal := 0

	for _, price := range prices {
		if price < minVal {
			minVal = price
		} else if price-minVal > maxVal {
			maxVal = price - minVal
		}
	}
	return maxVal
}
