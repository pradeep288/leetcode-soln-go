package medium_type

import "math"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := getDepth(root.Left)
	rightDepth := getDepth(root.Right)

	if leftDepth == rightDepth {
		return math.Pow(2, leftDepth)
	} else {
		return int(math.Pow(2, float32(rightDepth))) + getDepth(root.Left)
	}

}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + getDepth(root.Left)
}
