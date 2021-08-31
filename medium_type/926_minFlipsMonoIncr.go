package medium_type

func minFlipsMonoIncr(s string) int {
	var zero, one, output int

	for i:=0;i<len(s);i++{
		if s[i] == 48{
			zero += 1
		}
	}

	for i:=0;i<len(s);i++{
		if s[i] == 48{
			zero -=1
		}

		if s[i] == 49{
			one +=1
		}
		output = minM(output, zero+one)
	}

	return output
}

func minM(x int , y int)int{
	if x < y{
		return x
	}
	return y
}