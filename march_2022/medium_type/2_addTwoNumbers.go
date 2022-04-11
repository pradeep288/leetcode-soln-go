package medium_type

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res, curNode *ListNode
	var sum, carry int
	for l1 != nil && l2 != nil {
		sum = carry
		sum += l1.Val
		sum += l2.Val

		carry, sum = sum/10, sum%10
		newNode := &ListNode{
			Val:  sum,
			Next: nil,
		}
		if res == nil {
			res = newNode
			curNode = newNode
		} else {
			curNode.Next = newNode
			curNode = curNode.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		for l1 != nil {
			sum = carry + l1.Val
			carry, sum = sum/10, sum%10
			newNode := &ListNode{
				Val:  sum,
				Next: nil,
			}
			curNode.Next = newNode
			curNode = curNode.Next
			l1 = l1.Next
		}
	}

	if l2 != nil {
		for l2 != nil {
			sum = carry + l2.Val
			carry, sum = sum/10, sum%10
			newNode := &ListNode{
				Val:  sum,
				Next: nil,
			}
			curNode.Next = newNode
			curNode = curNode.Next
			l2 = l2.Next
		}
	}

	if carry > 0 {
		curNode.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return res
}
