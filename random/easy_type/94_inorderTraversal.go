package easy_type

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		res = append(res, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)
	return res
}
