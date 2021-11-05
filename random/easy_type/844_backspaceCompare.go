package easy_type

func backspaceCompare(s string, t string) bool {
	var str1, str2 []rune

	for _, val := range s {
		if val != '#' {
			str1 = append(str1, rune(val))
		} else {
			if len(str1) != 0 {
				str1 = str1[:len(str1)-1]
			}
		}
	}

	for _, val := range t {
		if val != '#' {
			str2 = append(str2, rune(val))
		} else {
			if len(str2) != 0 {
				str2 = str2[:len(str2)-1]
			}
		}
	}

	if len(str1) != len(str2) {
		return false
	}

	return string(str1) == string(str2)
}
