package medium_type

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 10*sum + root.Val
	}

	return helper(root.Left, 10*sum+root.Val) + helper(root.Right, 10*sum+root.Val)

}
