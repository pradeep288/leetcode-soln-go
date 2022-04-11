package medium_type

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if k == 0 {
		return head
	}
	// Get the length of the list
	var prevNode *ListNode
	curNode, nodeCount := head, 0
	for curNode != nil {
		nodeCount += 1
		prevNode = curNode
		curNode = curNode.Next
	}
	lastNode := prevNode

	// obtain splitIndex
	var splitIndex int
	if k > nodeCount {
		splitIndex = nodeCount - (k % nodeCount)
	} else {
		splitIndex = nodeCount - k
	}

	//rotate the list by breaking the list at splitIndex
	if splitIndex > 0 {
		curNode, nodeCount = head, 1
		for nodeCount < splitIndex {
			curNode = curNode.Next
			nodeCount += 1
		}

		lastNode.Next = head
		head = curNode.Next
		curNode.Next = nil
	}

	return head
}
