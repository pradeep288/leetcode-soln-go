package easy_type

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tempNode := &ListNode{
		Val:  0,
		Next: nil,
	}
	curNode := tempNode

	for true {
		if l1!= nil && l2 != nil {
			if l1.Val < l2.Val {
				curNode.Next = l1
				l1 = l1.Next
			} else {
				curNode.Next = l2
				l2 = l2.Next
			}
		} else {
			break
		}
		curNode = curNode.Next
	}

	if l1 == nil {
		curNode.Next = l2
	}
	if l2 == nil{
		curNode.Next = l1
	}

	return tempNode.Next
}
