package medium_type

type Node struct {
     Val int
     Children []*Node
}


func levelOrder(root *Node) [][]int {
	var res [][]int

	if root == nil{
		return res
	}

	var q []*Node

	q = append(q, root)

	for len(q)>0{
		size := len(q)
		var child []int
		for i:=0;i<size;i++{
			node := q[i]
			child =append(child, node.Val)
			for _,val:= range node.Children{
				q = append(q, val)
			}
		}
		res = append(res, child)
		q = q[size:]
	}
	return res
}
