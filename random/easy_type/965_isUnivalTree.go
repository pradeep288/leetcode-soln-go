package easy_type

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	var q []*TreeNode

	q = append(q, root)

	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			node := q[0]
			q = q[1:]
			if node.Val != root.Val {
				return false
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			size--
		}
	}
	return true
}
