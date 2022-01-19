package hard

func candy(ratings []int) int {
	var candies int
	left, right := []int{}, []int{}

	size := len(ratings)

	// Initialised both right and left with 1
	for i := 0; i < size; i++ {
		left = append(left, 1)
		right = append(right, 1)
	}

	// update the candy based on left childs's ratings
	for i := 1; i < size; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	// update the candy based on the right child's ratings
	for i := size - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}

	// get the sum
	for i := 0; i < size; i++ {
		candies += max(right[i], left[i])
	}

	return candies
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
