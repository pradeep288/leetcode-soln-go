package easy_type

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	var idx int
	for i := 0; i < len(t); i++ {
		if t[i] == s[idx] {
			idx++
		}
		if idx == len(s) {
			return true
		}
	}

	return idx == len(s)
}
