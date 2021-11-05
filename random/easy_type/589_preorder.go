package easy_type

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int
	var preOrderHelper func(*Node)
	preOrderHelper = func(root *Node) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		for _, child := range root.Children {
			preOrderHelper(child)
		}

	}
	preOrderHelper(root)
	return res
}
