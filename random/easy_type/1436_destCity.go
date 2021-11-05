package easy_type

func destCity(paths [][]string) string {
	graph := make(map[string]string)

	for _, path := range paths {
		graph[path[0]] = path[1]
	}
	curCity := paths[0][0]
	for graph[curCity] != "" {
		curCity = graph[curCity]
	}

	return curCity
}
