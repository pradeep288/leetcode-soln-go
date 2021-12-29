package medium_type

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
func connect(root *Node) *Node {
	var q []*Node

	if root == nil {
		return root
	}

	q = append(q, root)
	var levelOrder [][]*Node

	// Create a levelOrder List using BFS.
	for len(q) > 0 {
		size := len(q)
		var curLevelNodes []*Node
		for size > 0 {
			curNode := q[0]
			if curNode.Left != nil {
				q = append(q, curNode.Left)
			}
			if curNode.Right != nil {
				q = append(q, curNode.Right)
			}
			curLevelNodes = append(curLevelNodes, curNode)
			q = q[1:]
			size--
		}
		levelOrder = append(levelOrder, curLevelNodes)
	}

	// Update Node right using the level order
	for _, nodes := range levelOrder {
		for idx, node := range nodes {
			if idx == len(nodes)-1 {
				node.Next = nil
			} else {
				node.Next = nodes[idx+1]
			}
		}
	}

	return root
}*/

func connect(root *Node) *Node {
	var q []*Node

	if root == nil {
		return root
	}

	q = append(q, root)

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]

			if i != size-1 {
				cur.Next = q[0]
			}

			if cur.Left != nil {
				q = append(q, cur.Left)
			}

			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return root
}
