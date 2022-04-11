package easy_type

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth int
	q := []*TreeNode{root}

	for len(q) != 0 {
		size := len(q)
		depth++
		for size != 0 {
			curNode := q[0]
			q = q[1:]
			if curNode.Left != nil {
				q = append(q, curNode.Left)
			}
			if curNode.Right != nil {
				q = append(q, curNode.Right)
			}
			size--
		}
	}
	return depth
}*/

func maxDepth(root *TreeNode) int {
	var dfs func(root *TreeNode, depth int) int
	var max func(int, int) int
	max = func(i int, j int) int {
		if i > j {
			return i
		}
		return j
	}
	dfs = func(root *TreeNode, depth int) int {
		if root == nil {
			return depth
		}
		depth += 1

		return max(dfs(root.Left, depth), dfs(root.Right, depth))
	}
	return dfs(root, 0)
}
