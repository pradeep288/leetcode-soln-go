package medium_type

func kthSmallest(root *TreeNode, k int) int {
	var arr []int
	var traverseInOrder func(node *TreeNode)

	traverseInOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverseInOrder(node.Left)
		arr = append(arr, node.Val)
		traverseInOrder(node.Right)
	}
	traverseInOrder(root)

	return arr[k-1]
}
