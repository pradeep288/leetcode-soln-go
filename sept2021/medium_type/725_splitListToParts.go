package medium_type

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	// find the length of the list
	var totalLen int
	tmpHead := head

	for tmpHead != nil {
		totalLen += 1
		tmpHead = tmpHead.Next
	}

	len1 := totalLen / k
	distribute := totalLen % k

	var res []*ListNode
	cur := head
	i := totalLen

	for k > 0 {
		if i == 0 {
			res = append(res, nil)
		} else {
			count := len1
			if distribute > 0 {
				count += 1
				distribute -= 1
			}
			i -= count
			tmp, prev := cur, cur
			for count > 0 {
				prev = cur
				cur = cur.Next
				count -= 1
			}

			prev.Next = nil
			res = append(res, tmp)
		}
		k -= 1
	}
	return res

}
