package medium_type

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}

	curNode := root

	for true {
		if curNode.Val > val {
			if curNode.Left == nil {
				curNode.Left = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}
				return root
			}
			curNode = curNode.Left
		} else if curNode.Val < val {
			if curNode.Right == nil {
				curNode.Right = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}
				return root
			}
			curNode = curNode.Right
		}
	}

	return root
}

/*

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil{
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}

	if root.Val<val{
		root.Right = insertIntoBST(root.Right, val)
	}

	if root.Val<val{
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}
*/
