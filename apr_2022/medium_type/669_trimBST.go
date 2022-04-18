package medium_type

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var arr []*TreeNode
	var preOrderTraversal func(node *TreeNode)

	preOrderTraversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		arr = append(arr, node)
		preOrderTraversal(node.Left)
		preOrderTraversal(node.Right)
	}

	preOrderTraversal(root)
	var head *TreeNode

	// [3,0,4,null,2,null,null,1]
	for i := 0; i < len(arr); i++ {
		if arr[i].Val < low || arr[i].Val > high {
			continue
		}
		newNode := &TreeNode{
			Val:   arr[i].Val,
			Left:  nil,
			Right: nil,
		}
		if head == nil {
			head = newNode
		} else {
			curNode := head
			var prevNode *TreeNode
			for curNode != nil {
				prevNode = curNode
				if curNode.Val > arr[i].Val {
					curNode = curNode.Left
				} else {
					curNode = curNode.Right
				}
			}
			if prevNode.Val > arr[i].Val {
				prevNode.Left = newNode
			} else {
				prevNode.Right = newNode
			}
		}
	}

	return head
}
