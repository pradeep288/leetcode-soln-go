package easy_type

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	// Checks Boundary condition
	if root == nil || root.Left == nil && root.Right == nil {
		return 0
	}
	return helper(root, nil, 0)
}

func helper(cur *TreeNode, parent *TreeNode, sum int) int {
	if cur == nil {
		return 0
	}

	// Check the given node(cur) is leaf node and node is left of its parent
	if cur.Left == nil && cur.Right == nil && parent.Left == cur {
		sum += cur.Val
		return sum
	}

	rSum := helper(cur.Right, cur, sum)
	lSum := helper(cur.Left, cur, sum)

	return lSum + rSum
}
