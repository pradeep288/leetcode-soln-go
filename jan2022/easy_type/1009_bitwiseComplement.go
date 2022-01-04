package easy_type

import "math"

/*
func bitwiseComplement(n int) int {
	var res, i int
	i = 1
	for n > 0 {
		if n&1 == 0 {
			res += 1 << i
		}
		n >>= 1
		i++
	}
	return res
}*/

func bitwiseComplement(n int) int {
	var res int
	var complement []int
	// handle corner case
	if n == 0 {
		return 1
	}

	// update complement slice with complement of each digit in reverse order
	for n > 0 {
		if n&1 == 0 {
			complement = append(complement, 1)
		} else {
			complement = append(complement, 0)
		}
		n >>= 1
	}

	// obtain the result in decimal format
	for i := 0; i < len(complement); i++ {
		res += int(math.Pow(float64(2), float64(i)) * float64(complement[i]))
	}
	return res
}
