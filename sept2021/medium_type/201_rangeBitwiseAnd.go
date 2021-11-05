package medium_type

// https://www.youtube.com/watch?v=-qrpJykY2gE&ab_channel=TECHDOSE

func rangeBitwiseAnd(left int, right int) int {
	var count int

	for left != right {
		left >>= 1
		right >>= 1
		count += 1
	}

	return left << count
}
