package easy_type

func minDepth(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := dfs(root.Left)
	rightDepth := dfs(root.Right)

	if leftDepth == 0 || rightDepth == 0 {
		return leftDepth + rightDepth + 1
	}
	return min(leftDepth, rightDepth) + 1

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
