package medium_type

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	var q []*TreeNode
	var maxDiff int
	maxDiff = math.MinInt32
	if root == nil {
		return 0
	}

	q = append(q, root)

	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			curNode := q[0]
			diff := getSubtreeDiff(curNode)
			if maxDiff < diff {
				maxDiff = diff
			}
			if curNode.Left != nil {
				q = append(q, curNode.Left)
			}
			if curNode.Right != nil {
				q = append(q, curNode.Right)
			}
			q = q[1:]
			size--
		}
	}
	return maxDiff
}

func getSubtreeDiff(root *TreeNode) int {
	var q []*TreeNode
	var maxCurSubtreeDiff int
	maxCurSubtreeDiff = math.MinInt32

	q = append(q, root)

	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			curNode := q[0]
			d := getAbsDiff(root.Val, curNode.Val)
			if d > maxCurSubtreeDiff {
				maxCurSubtreeDiff = d
			}
			if curNode.Left != nil {
				q = append(q, curNode.Left)
			}
			if curNode.Right != nil {
				q = append(q, curNode.Right)
			}
			q = q[1:]
			size--
		}
	}
	return maxCurSubtreeDiff
}

func getAbsDiff(a, b int) int {
	d := a - b
	if d < 0 {
		return d * -1
	}
	return d
}
