package medium_type

func searchMatrix(matrix [][]int, target int) bool {
	binarySearch := func(arr []int, target int) bool {
		left, right := 0, len(arr)-1
		for left <= right {
			mid := left + (right-left)/2
			if arr[mid] == target {
				return true
			}
			if arr[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return false
	}

	twoDToOneDArr := make([]int, len(matrix)*len(matrix[0]))
	var count int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			twoDToOneDArr[count] = matrix[i][j]
			count++
		}
	}

	return binarySearch(twoDToOneDArr, target)
}
