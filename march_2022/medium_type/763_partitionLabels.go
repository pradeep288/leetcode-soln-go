package medium_type

// https://www.youtube.com/watch?v=B7m8UmZE-vw
func partitionLabels(s string) []int {
	lastIdx := make(map[string]int)

	for i, c := range s {
		lastIdx[string(c)] = i
	}

	getMax := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	var res []int
	size, end := 0, 0
	for i, c := range s {
		size += 1
		end = getMax(end, lastIdx[string(c)])
		if size == i {
			res = append(res, size)
			size = 0
		}
	}

	return res
}
