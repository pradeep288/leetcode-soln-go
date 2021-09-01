package easy_type

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0{
		return nil
	}
	return constructBST(nums, 0, len(nums)-1)
}

func constructBST(nums []int, left int, right int) *TreeNode{
	if left> right{
		return nil
	}

	mid := (left + (right-left) )/2

	node:= TreeNode{Val: nums[mid]}
	node.Left = constructBST(nums, left, mid -1)
	node.Right = constructBST(nums, mid +1, right)

	return &node
}