package easy_type

import "math"

/*
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n = n / 2
	}
	return true
}


func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	var count int
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	if count == 1 {
		return true
	}
	return false
}
*/
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	res := math.Log2(float64(n))
	if math.Ceil(res) == math.Floor(res) {
		return true
	}
	return false
}
