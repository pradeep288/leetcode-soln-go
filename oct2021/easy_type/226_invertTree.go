// https://www.youtube.com/watch?v=_i0jqdVkObU&ab_channel=TECHDOSE
package easy_type

func invertTree(root *TreeNode) *TreeNode {
	return helper(root)
}

func helper(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := helper(root.Left)
	right := helper(root.Right)

	root.Left, root.Right = right, left

	return root
}
