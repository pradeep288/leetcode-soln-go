package medium_type

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur, nextNode := head, head.Next

	for true {
		cur.Val, nextNode.Val = nextNode.Val, cur.Val
		if nextNode.Next == nil || nextNode.Next.Next == nil {
			break
		}
		cur = nextNode.Next
		nextNode = nextNode.Next.Next
	}
	return head
}
