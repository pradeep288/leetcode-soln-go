// Uses Greedy Approach, BFS on the 1D array.
// Refenece Used: https://www.youtube.com/watch?v=dJ7sWiOoK7g
func jump(nums []int) int {
	left, right, res := 0, 0, 0

	for right < len(nums)-1{
		farthest := 0
		for i:=left; i<=right;i++{
			farthest = max(farthest, i + nums[i])
		}
		left = right + 1
		right = farthest
		res += 1
	}
	return res
}

func max(x int, y int) int{
	if x>y {
		return x
	}
	return y
}
