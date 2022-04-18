package easy_type

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	curNode := root

	for curNode != nil {
		if curNode.Val == val {
			return curNode
		}
		if curNode.Val < val {
			curNode = curNode.Right
		} else {
			curNode = curNode.Left
		}
	}
	return nil
}
