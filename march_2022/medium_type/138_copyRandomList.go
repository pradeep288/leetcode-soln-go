package medium_type

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	ptr1 := head
	for ptr1 != nil {
		newNode := &Node{
			Val:    ptr1.Val,
			Next:   ptr1.Next,
			Random: ptr1.Random,
		}
		ptr1.Next = newNode
		ptr1 = newNode.Next
	}

	// Update the random pointer.
	ptr1 = head
	for ptr1 != nil {
		ptr2 := ptr1.Next
		if ptr1.Random == nil {
			ptr2.Random = nil
		} else {
			ptr2.Random = ptr1.Random.Next
		}
		ptr1 = ptr1.Next.Next
	}

	ptr1 = head
	newHead := head.Next
	for ptr1 != nil {
		ptr2 := ptr1.Next
		ptr1.Next = ptr2.Next
		if ptr1.Next != nil {
			ptr2.Next = ptr1.Next.Next
		} else {
			ptr2.Next = nil
		}
		ptr1 = ptr1.Next
	}

	return newHead
}
