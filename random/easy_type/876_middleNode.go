package easy_type

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	sp, fp := head, head
	for fp.Next != nil {
		fp = fp.Next.Next
		sp = sp.Next
	}
	return sp
}
