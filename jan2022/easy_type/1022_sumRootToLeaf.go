package easy_type

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {

	var dfs func(node *TreeNode, curVal int) int

	dfs = func(root *TreeNode, curVal int) int {
		if root == nil {
			return 0
		}
		curVal <<= 1
		curVal += root.Val
		if root.Left == nil && root.Right == nil {
			return curVal
		}
		return dfs(root.Left, curVal) + dfs(root.Right, curVal)
	}

	return dfs(root, 0)
}
