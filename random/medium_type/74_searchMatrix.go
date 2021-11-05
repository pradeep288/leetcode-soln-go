// https://www.youtube.com/watch?v=ZYpYur0znng&ab_channel=takeUforward
package medium_type

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	low, high := 0, (rows*cols)-1

	for low <= high {
		mid := (low + (high-low)/2)
		if matrix[mid/cols][mid%cols] == target {
			return true
		}

		if matrix[mid/cols][mid%cols] < target {
			low += 1
		} else {
			high -= 1
		}
	}
	return false
}
