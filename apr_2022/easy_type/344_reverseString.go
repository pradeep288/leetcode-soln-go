package easy_type

func reverseString(s []byte) {
	sp, ep := 0, len(s)-1

	for sp < ep {
		s[sp], s[ep] = s[ep], s[sp]
		sp += 1
		ep -= 1
	}
}
