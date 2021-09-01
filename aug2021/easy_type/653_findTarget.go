package easy_type

var visited []int
func findTarget(root *TreeNode, k int) bool {
	visited = []int{}
	return dfsTarget(root, k)
}

func dfsTarget(root *TreeNode, k int)  bool{
	if root == nil{
		return false
	}
	diff := k - root.Val

	if valExists(diff, visited){
		return true
	}
	visited = append(visited, root.Val)

	return dfsTarget(root.Left, k) || dfsTarget(root.Right , k)
}

func valExists(target int, visited []int) bool{
	for _,v:=range visited{
		if target==v{
			return true
		}
	}
	return false
}