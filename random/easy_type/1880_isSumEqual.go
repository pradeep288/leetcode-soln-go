package easy_type

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	firstWordVal := wordValue(firstWord)
	secondWordVal := wordValue(secondWord)
	targetWordVal := wordValue(targetWord)

	if firstWordVal+secondWordVal == targetWordVal {
		return true
	}
	return false
}

func wordValue(word string) rune {
	var val, mul rune
	mul = 1
	for i := len(word) - 1; i >= 0; i-- {
		v := rune(word[i]) - 'a'
		val += v * mul
		mul *= 10
	}
	return val
}
