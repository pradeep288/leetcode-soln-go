package easy_type

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	sp, fp := head, head

	for fp.Next != nil && fp.Next.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
	}
	// Return the second middle element when length of the list is odd
	if fp.Next == nil {
		return sp
	}

	// Return the middle element when length of the list is even
	return sp.Next
}
