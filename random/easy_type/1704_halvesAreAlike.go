package easy_type

func halvesAreAlike(s string) bool {
	var half1, half2, mid int
	mid = len(s) / 2

	for i, c := range s {
		if i < mid {
			switch c {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				half1++
			}
		} else {
			switch c {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				half2++
			}
		}
	}
	return half1 == half2
}
