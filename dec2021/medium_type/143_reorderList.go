package medium_type

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	var stack []*ListNode

	curNode := head

	// Update the stack with Nodes
	for curNode != nil {
		stack = append(stack, curNode)
		curNode = curNode.Next
	}

	if len(stack) == 1 {
		curNode = stack[0]
		stack = stack[:len(stack)-1]
	} else {
		// Reorder the head and its next node
		head = stack[0]
		stack = stack[1:]
		head.Next = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curNode = head.Next

		// Reorder the remaining elements
		for len(stack) >= 2 {
			node1 := stack[0]
			curNode.Next = node1
			curNode = node1
			stack = stack[1:]
			node2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curNode.Next = node2
			curNode = node2
		}
		// Handle when the length of list is Odd
		if len(stack) == 1 {
			lastNode := stack[0]
			curNode.Next = lastNode
			curNode = lastNode
		}
	}
	// Update the last element
	curNode.Next = nil
}
