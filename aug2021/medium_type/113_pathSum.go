package medium_type


type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	path:=make([]int,0)
	if root == nil{
		return res
	}
	dfs(root, 0, path,  targetSum, &res)
	return res
}

func dfs(node *TreeNode, sum int, path []int, targetSum int, res *[][]int){
	sum += node.Val
	path = append(path, node.Val)
	if node.Left == nil && node.Right == nil && sum == targetSum {
		cpy := make([]int, len(path))
		copy(cpy, path)
		*res = append(*res, cpy)
	}

	if node.Left != nil{
		dfs(node.Left, sum , path, targetSum, res)
	}
	if node.Right!= nil{
		dfs(node.Right, sum, path, targetSum, res)
	}
}
