package medium_type

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	// Handling the Edge case
	if head == nil || head.Next == nil {
		return head
	}

	oddHead, curOdd := head, head
	evenHead, curEven := head.Next, head.Next
	curNode := curEven.Next
	count := 1
	for curNode != nil {
		if count%2 == 0 {
			curEven.Next = curNode
			curEven = curEven.Next
		} else {
			curOdd.Next = curNode
			curOdd = curOdd.Next
		}
		curNode = curNode.Next
		count += 1
	}
	if count%2 == 0 {
		curEven.Next = nil
	}
	curOdd.Next = evenHead
	return oddHead
}
