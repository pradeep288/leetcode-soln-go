package easy_type

func validPath(n int, edges [][]int, start int, end int) bool {
	if n == 1 {
		return true
	}

	path := make(map[int][]int)
	for _, edge := range edges {
		path[edge[0]] = append(path[edge[0]], edge[1])
		path[edge[1]] = append(path[edge[1]], edge[0])

	}

	q := [][]int{path[start]}
	visited := make(map[int]interface{})
	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			nodes := q[0]
			for _, node := range nodes {
				if _, ok := visited[node]; !ok {
					if node == end {
						return true
					}
					visited[node] = true
					q = append(q, path[node])

				}
			}
			q = q[1:]
			size--
		}

	}

	return false

}
