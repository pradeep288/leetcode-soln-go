package medium_type

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	var list []*TreeNode

	if start > end {
		list = append(list, nil)
		return list
	}

	if start == end {
		list = append(list, &TreeNode{
			Val: start,
		})
		return list
	}

	for i := start; i <= end; i++ {
		leftPossibleTrees := helper(start, i-1)
		rightPossibleTrees := helper(i+1, end)
		for _, l := range leftPossibleTrees {
			for _, r := range rightPossibleTrees {
				root := TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				}
				list = append(list, &root)
			}
		}
	}
	return list
}
