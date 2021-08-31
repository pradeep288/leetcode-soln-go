package medium_type

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/
var numOfGoodNodes int

func goodNodes(root *TreeNode) int {
	numOfGoodNodes=0
	numOfGoodNodesDFS(root, -100000)
	return numOfGoodNodes
}

func numOfGoodNodesDFS(root *TreeNode, curMax int)  {
	if root == nil{
		return
	}
	if root.Val>=curMax{
		curMax = root.Val
		numOfGoodNodes += 1
	}
	numOfGoodNodesDFS(root.Right, curMax)
	numOfGoodNodesDFS(root.Left, curMax)
}