package hard_type

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	var headSet bool

	start := head

	var lenOfList int
	for start != nil {
		start = start.Next
		lenOfList++
	}

	var cur, prev, next *ListNode
	numOfGroups := lenOfList / k
	cur = head

	var hashMap [][]*ListNode
	for numOfGroups > 0 {
		count := k
		prev = &ListNode{}
		firstNode := cur
		for count > 0 {
			next = cur.Next
			cur.Next = prev
			prev = cur
			cur = next
			count--
		}
		if !headSet {
			head = prev
			headSet = true
		}

		numOfGroups--
		lastNode := prev
		hashMap = append(hashMap, []*ListNode{firstNode, lastNode})
	}

	for i := 0; i < len(hashMap)-1; i++ {
		hashMap[i][0].Next = hashMap[i+1][1]
	}
	if lenOfList%k == 0 {
		hashMap[len(hashMap)-1][0].Next = nil
	} else {
		hashMap[len(hashMap)-1][0].Next = next
	}

	return head
}
