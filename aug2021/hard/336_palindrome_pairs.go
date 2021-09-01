package hard

func palindromePairs(words []string) [][]int {
	var res [][]int

	// Create a map of value and index
	var inputMap = make(map[string]int)
	for i, val := range words {
		inputMap[val] = i
	}

	for i, v := range words {
		// empty string + Palindrome and palindrome + ""
		if v == ""{
			for j:=0;j<=len(words)-1;j++{
				if isPalindrome(words[j]) && i!=j {
					res = append(res, []int{i,j})
					res = append(res, []int{j,i})
				}
			}
		}

		// string + reverse of string
		rev := reverse(words[i])
		if j,ok := inputMap[rev];ok && i!=j{
			res = append(res,[]int{i,j})
		}

		// 3rd case
		// palindrome + x and rev(x) exists
		// x + palindrome and rev(x) exists
		for k:=1;k<=len(words[i])-1;k++{
			rev = reverse(v[k:])
			if _,ok:=inputMap[rev];ok && isPalindrome(v[:k]) && i != inputMap[rev] {
				res = append(res,[]int{inputMap[rev], i})
			}

			rev = reverse(v[:k])
			if _,ok:=inputMap[rev];ok && isPalindrome(v[k:]) && i != inputMap[rev] {
				res = append(res,[]int{i, inputMap[rev]})
			}
		}
	}
	return res
}

func isPalindrome(word string) bool {
	j := len(word) - 1
	for i := 0; i <= j; i++ {
		if !(word[i] == word[j]) {
			return false
		}
		j -= 1
	}
	return true
}


func reverse(inp string) string {
	val := []rune(inp)
	var out string
	for i:=len(val)-1;i>=0;i--{
		out = out + string(val[i])
	}
	return out
}

