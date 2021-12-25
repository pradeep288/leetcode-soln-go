package medium_type

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Create a Graph from the input
	graph := make(map[int][]int)

	for _, edge := range prerequisites {
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	visited := make([]bool, len(graph))
	for i := 0; i < len(graph); i++ {
		visited[i] = true
		for j := 0; j < len(graph[i]); j++ {
			FLAG := isCyclicUtil(graph, visited, graph[i][j])
			if FLAG == true {
				return false
			}
		}
		visited[i] = false
	}

	return true
}

func isCyclicUtil(graph map[int][]int, visited []bool, node int) bool {
	if visited[node] == true {
		return true
	}
	visited[node] = true

	for i := 0; i < len(graph[node]); i++ {
		FLAG := isCyclicUtil(graph, visited, graph[node][i])
		if FLAG == true {
			return true
		}
	}
	return false
}
