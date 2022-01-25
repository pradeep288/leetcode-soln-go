package easy_type

func validMountainArray(arr []int) bool {
	var mountainExists, valleyExists bool
	var valleyStart, valleyEnd int

	for i := 0; i <= len(arr)-2; i++ {
		if arr[i] < arr[i+1] {
			mountainExists = true
		} else {
			valleyStart = i
			break
		}
	}

	for i := valleyStart; i <= len(arr)-2; i++ {
		if arr[i] > arr[i+1] {
			valleyExists = true
			valleyEnd = i
		} else {
			break
		}
	}

	return mountainExists && valleyExists && valleyEnd+1 == len(arr)-1
}
