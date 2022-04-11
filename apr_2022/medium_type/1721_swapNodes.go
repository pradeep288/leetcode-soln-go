package medium_type

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	// find the length of List
	var count int
	curNode := head
	for curNode != nil {
		count++
		curNode = curNode.Next
	}

	// identify k th node last
	i := 0
	curNode = head
	for i < count-k {
		i++
		curNode = curNode.Next
	}
	endNode := curNode

	// identify k th node from the first
	i = 1
	curNode = head
	for i < k {
		i++
		curNode = curNode.Next
	}
	firstNode := curNode

	// swap node values
	firstNode.Val, endNode.Val = endNode.Val, firstNode.Val

	return head
}
