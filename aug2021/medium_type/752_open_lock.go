package medium_type

// Use BFS Method to solve the problem
func openLock(deadends []string, target string) int {
	var queue []string

	queue = append(queue, "0000")


	deads := make(map[string]struct{})
	for _, dead := range deadends {
		deads[dead] = struct{}{}
	}

	visited := make(map[string]struct{})

	level:=0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			newCode := queue[0]
			queue = queue[1:]

			if _, ok := deads[newCode]; ok {
				continue
			}

			if _,ok:=visited[newCode];ok {
				continue
			}

			if newCode == target {
				return level
			}

			for i, v := range newCode {
				x := incrementVal(v)
				s1 := newCode[:i] + string(x) + newCode[i+1:]

				_, x1 := deads[s1]
				_, x2 := visited[s1]
				if !x1 && !x2 {
					queue = append(queue, s1)
				}

				y := decrementVal(v)
				s2 := newCode[:i] + string(y) + newCode[i+1:]

				_, y1 := deads[s2]
				_, y2 := visited[s2]
				if !y1 && !y2 {
					queue = append(queue, s2)
				}
			}
			visited[newCode] = struct{}{}
		}
		level += 1
	}
	return -1
}



func incrementVal(x int32) int32 {
	if x == 57 {
		return 48
	}
	return x+1
}


func decrementVal(x int32) int32 {
	if x == 48 {
		return 57
	}
	return x-1
}

