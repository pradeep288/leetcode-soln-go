package easy_type

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, temp *ListNode
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		head, temp = list1, list1
		list1 = list1.Next
	} else {
		head, temp = list2, list2
		list2 = list2.Next
	}

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			temp.Next = list1
			temp = temp.Next
			list1 = list1.Next
		} else {
			temp.Next = list2
			temp = temp.Next
			list2 = list2.Next
		}
	}

	if list1 != nil {
		temp.Next = list1
	}
	if list2 != nil {
		temp.Next = list2
	}

	return head
}
