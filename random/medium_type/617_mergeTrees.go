// https://www.youtube.com/watch?v=lw5swOzpFtE&t=38s
package medium_type

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return root1
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	var q [][]*TreeNode

	q = append(q, []*TreeNode{root1, root2})

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if cur[1] != nil {
			cur[0].Val += cur[1].Val

			if cur[0].Left == nil {
				cur[0].Left = cur[1].Left
			} else {
				q = append(q, []*TreeNode{cur[0].Left, cur[1].Left})
			}

			if cur[0].Right == nil {
				cur[0].Right = cur[1].Right
			} else {
				q = append(q, []*TreeNode{cur[0].Right, cur[1].Right})
			}
		}
	}
	return root1
}
