package medium_type

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var items []int

	var dfs func(*TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		items = append(items, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root1)
	dfs(root2)

	sort.Ints(items)

	return items

}
