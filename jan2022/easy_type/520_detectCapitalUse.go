package easy_type

func detectCapitalUse(word string) bool {
	if allCapital(word) {
		return true
	}

	if allSmall(word) {
		return true
	}

	if capitalized(word) {
		return true
	}

	return false
}

func allCapital(word string) bool {
	for _, c := range word {
		if !(c >= 'A' && c <= 'Z') {
			return false
		}
	}
	return true
}

func allSmall(word string) bool {
	for _, c := range word {
		if !(c >= 'a' && c <= 'z') {
			return false
		}
	}
	return true
}

func capitalized(word string) bool {
	for idx, c := range word {
		if idx == 0 && !(c >= 'A' && c <= 'Z') {
			return false
		}
		if idx > 0 && !(c >= 'a' && c <= 'z') {
			return false
		}
	}
	return true
}
