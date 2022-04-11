package easy_type

func checkValid(matrix [][]int) bool {
	var resetMap func(int) map[int]bool

	resetMap = func(len int) map[int]bool {
		m := make(map[int]bool)
		for i := 0; i < len; i++ {
			m[i+1] = false
		}
		return m
	}

	for i := 0; i < len(matrix); i++ {
		m := resetMap(len(matrix))
		for j := 0; j < len(matrix); j++ {
			if m[matrix[i][j]] {
				return false
			} else {
				m[matrix[i][j]] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		m := resetMap(len(matrix))
		for j := 0; j < len(matrix); j++ {
			if m[matrix[j][i]] {
				return false
			} else {
				m[matrix[j][i]] = true
			}
		}
	}

	return true
}
