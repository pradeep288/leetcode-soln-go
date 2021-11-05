package medium_type

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}

	var stack []*Node
	stack = append(stack, root)

	prev := &Node{
		Val:  0,
		Prev: nil,
		Next: nil,
	}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur.Prev = prev
		prev.Next = cur
		prev = cur
		if cur.Next != nil {
			stack = append(stack, cur.Next)
		}
		if cur.Child != nil {
			stack = append(stack, cur.Child)
			cur.Child = nil
		}
	}

	root.Prev = nil
	return root
}
