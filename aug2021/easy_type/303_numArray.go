package easy_type

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	var res int
	for left<=right{
		res += this.nums[left]
		left+=1
	}
	return res
}
