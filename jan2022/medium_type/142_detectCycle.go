package medium_type

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
	sp, fp := head, head
	for fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
		if sp == fp {
			break
		}
	}
	if fp == nil || fp.Next == nil {
		return nil
	}
	fp = head

	for fp != sp {
		fp = fp.Next
		sp = sp.Next
	}
	return fp
}
