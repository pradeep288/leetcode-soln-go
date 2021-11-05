package medium_type

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func pathSum(root *TreeNode, targetSum int) int {
	ans = 0
	baseHelper(root, targetSum)
	return ans
}
func baseHelper(root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	helper(root, targetSum)
	baseHelper(root.Left, targetSum)
	baseHelper(root.Right, targetSum)
}

func helper(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	if sum == root.Val {
		ans++
	}
	sum -= root.Val
	helper(root.Left, sum)
	helper(root.Right, sum)
}
