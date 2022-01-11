package medium_type

import (
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	node *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		node: head,
	}
}

func (this *Solution) GetRandom() int {
	cur := this.node
	res := cur.Val
	count := 1

	for cur != nil {
		v := rand.Intn(count)
		if v == 0 {
			res = cur.Val
		}
		cur = cur.Next
		count = count + 1
	}

	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
