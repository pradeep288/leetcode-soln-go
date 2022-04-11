package medium_type

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type info struct {
		node  *TreeNode
		index int
	}

	var q []info
	maxWidth := math.MinInt32
	q = append(q, info{
		node:  root,
		index: 0,
	})

	for len(q) > 0 {
		curWidth := q[len(q)-1].index - q[0].index + 1
		if curWidth > maxWidth {
			maxWidth = curWidth
		}
		size := len(q)
		for size > 0 {
			curNode := q[0]
			q = q[1:]
			size -= 1
			if curNode.node.Left != nil {
				q = append(q, info{
					node:  curNode.node.Left,
					index: 2 * curNode.index,
				})
			}
			if curNode.node.Right != nil {
				q = append(q, info{
					node:  curNode.node.Right,
					index: 2*curNode.index + 1,
				})
			}
		}
	}

	return maxWidth
}
