package medium_type

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead, temp *ListNode
	prev, cur, next := &ListNode{Val: -101, Next: head}, head, head.Next
	for cur != nil && next != nil {
		if prev.Val != cur.Val && cur.Val != next.Val {
			newNode := &ListNode{
				cur.Val,
				nil,
			}
			if newHead == nil {
				newHead = newNode
				temp = newHead

			} else {
				temp.Next = newNode
				temp = temp.Next
			}
		}
		cur = cur.Next
		prev = prev.Next
		next = next.Next
	}

	if cur.Val != prev.Val {
		newNode := &ListNode{
			cur.Val,
			nil,
		}
		if newHead == nil {
			return newNode
		}
		temp.Next = newNode
	}
	return newHead
}
