package medium_type

import "strings"

func shiftingLetters(s string, shifts []int) string {
	var out []string
	t := totalShifts(shifts)
	for i, v := range s {
		if i != 0 {
			t = t - int64(shifts[i-1])
		}
		out = append(out, string(shift(int64(v), t)))
	}
	return strings.Join(out, "")
}

func shift(v int64, s int64) int64 {
	var res int64
	rem := s % 26

	if rem+v > 122 {
		res = 96 + rem + v - 122
	} else {
		res = rem + v
	}
	return res
}

func totalShifts(s []int) int64 {
	var count int64
	for _, v := range s {
		count += int64(v)
	}
	return count
}
