package easy_type

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return head
	}

	var prev, cur, next *ListNode
	cur = head
	next = cur.Next

	for next != nil{
		cur.Next = prev
		prev = cur
		cur = next
		next = next.Next
	}
	cur.Next = prev
	return cur
}*/
func reverseList(head *ListNode) *ListNode {
	return helper(nil, head)
}

func helper(prev *ListNode, cur *ListNode) *ListNode {
	if cur == nil {
		return prev
	}
	tail := helper(cur, cur.Next)
	cur.Next = prev
	return tail
}
