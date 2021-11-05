package easy_type

func validMountainArray(arr []int) bool {
	var valleyStart int
	var mountain, valley bool

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			valleyStart = i
			break
		}
		mountain = true
	}

	for i := valleyStart; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			return false
		}
		valley = true
	}
	if mountain && valley {
		return true
	}

	return false
}
