package easy_type

func findSpecialInteger(arr []int) int {
	n := len(arr)
	t := n / 4

	for i := 0; i < n-t; i++ {
		if arr[i] == arr[i+t] {
			return arr[i]
		}
	}

	return -1
}
