package easy_type

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {

	treeMap := createTreeMap(root)
	if treeMap[x][0] != treeMap[y][0] && treeMap[x][1] == treeMap[y][1] {
		return true
	}
	return false

}

func createTreeMap(root *TreeNode) map[int][]int {
	treeMap := make(map[int][]int)
	if root == nil {
		return treeMap
	}

	var q []*TreeNode
	q = append(q, root)
	var depth int

	treeMap[root.Val] = append(treeMap[root.Val], -1, depth)

	for len(q) != 0 {
		depth += 1
		size := len(q)
		for size > 0 {
			node := q[0]

			if node.Left != nil {
				treeMap[node.Left.Val] = append(treeMap[node.Left.Val], node.Val, depth)
				q = append(q, node.Left)
			}
			if node.Right != nil {
				treeMap[node.Right.Val] = append(treeMap[node.Right.Val], node.Val, depth)
				q = append(q, node.Right)
			}
			q = q[1:]
			size--
		}

	}
	return treeMap
}
