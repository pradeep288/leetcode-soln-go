package easy_type

func increasingBST(root *TreeNode) *TreeNode {
	var arr []int
	var inOrderTraversal func(node *TreeNode)

	inOrderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrderTraversal(root.Left)
		arr = append(arr, root.Val)
		inOrderTraversal(root.Right)
	}

	// Get the values of nodes in increasing order by Inorder traversal
	inOrderTraversal(root)

	// Recreate the BST using the values obtained in the earlier step
	head := &TreeNode{}
	cur := head
	for i := 0; i < len(arr); i++ {
		newNode := &TreeNode{
			Val:   arr[i],
			Left:  nil,
			Right: nil,
		}
		cur.Right = newNode
		cur = newNode
	}

	return head.Right
}
