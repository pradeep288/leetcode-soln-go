package easy_type

func hasCycle(head *ListNode) bool {
	var fp, sp *ListNode
	sp = head
	fp = head.Next.Next

	for sp != nil && fp != nil {
		if sp == fp {
			return true
		}
		sp = sp.Next
		fp = fp.Next.Next
	}

	return false
}
