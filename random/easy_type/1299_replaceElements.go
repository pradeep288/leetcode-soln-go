package easy_type

func replaceElements(arr []int) []int {
	max_val := arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = max_val
		if max_val < tmp {
			max_val = tmp
		}
	}
	arr[len(arr)-1] = -1
	return arr
}
