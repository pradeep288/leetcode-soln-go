package easy_type

func validPalindrome(s string) bool {
	isPalindrome := func(s string, leftBound, rightBound int) bool {
		for left, right := leftBound, rightBound; left < right; left, right = left+1, right-1 {
			if s[left] != s[right] {
				return false
			}
		}
		return true
	}

	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
	}

	return true
}
