package easy_type

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	sp, fp := head, head.Next

	for fp != nil && fp.Next != nil {
		if sp == fp {
			return true
		}
		sp = sp.Next
		fp = fp.Next.Next
	}
	return false
}
