package hard_type

//Uses Jarvis March Algorithm.
// Please refer this video https://www.youtube.com/watch?v=Vu84lmMzP2o&t=128s&ab_channel=TusharRoy-CodingMadeSimple
func outerTrees(trees [][]int) [][]int {
	// Identify the leftmost tree.
	start := trees[0]
	for _, p := range trees {
		if start[0] > p[0] {
			start = p
		}
	}

	current := start
	var res [][]int
	res = append(res, current)

	var coLinearPoints [][]int

	for true {
		nextTarget := trees[0]

		for i := 1; i < len(trees); i++ {
			if trees[i][0] == current[0] && trees[i][1] == current[1] {
				continue
			}
			val := crossProduct(current, nextTarget, trees[i])
			if val > 0 {
				nextTarget = trees[i]
				coLinearPoints = [][]int{}
			} else if val == 0 {
				if distance(current, nextTarget, trees[i]) < 0 {
					coLinearPoints = append(coLinearPoints, nextTarget)
					nextTarget = trees[i]
				} else {
					coLinearPoints = append(coLinearPoints, trees[i])
				}
			}
		}
		for _, p := range coLinearPoints {

			if !alreadyExists(res, p) {
				res = append(res, p)
			}
		}
		if nextTarget[0] == start[0] && nextTarget[1] == start[1] {
			break
		}
		if !alreadyExists(res, nextTarget) {
			res = append(res, nextTarget)
		}
		current = nextTarget
	}
	return res
}

func crossProduct(a, b, c []int) int {
	y1 := a[1] - b[1]
	y2 := a[1] - c[1]
	x1 := a[0] - b[0]
	x2 := a[0] - c[0]

	return y2*x1 - y1*x2
}

func distance(a, b, c []int) int {
	y1 := a[1] - b[1]
	y2 := a[1] - c[1]
	x1 := a[0] - b[0]
	x2 := a[0] - c[0]

	x := y1*y1 + x1*x1
	y := y2*y2 + x2*x2

	if x == y {
		return 0
	}
	if x < y {
		return -1
	}
	return 1
}

func alreadyExists(in [][]int, t []int) bool {
	for _, x := range in {
		if x[0] == t[0] && x[1] == t[1] {
			return true
		}
	}
	return false
}
