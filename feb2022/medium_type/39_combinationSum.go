package medium_type

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var backtracking func([]int, int, int)
	backtracking = func(curPath []int, curIndx int, total int) {
		if total == target {
			temp := make([]int, len(curPath))
			_ = copy(temp, curPath)
			res = append(res, temp)
			return
		}

		if curIndx >= len(candidates) || total > target {
			return
		}

		curPath = append(curPath, candidates[curIndx])
		backtracking(curPath, curIndx, total+candidates[curIndx])
		curPath = curPath[:len(curPath)-1]
		backtracking(curPath, curIndx+1, total)
	}

	backtracking([]int{}, 0, 0)

	return res
}
