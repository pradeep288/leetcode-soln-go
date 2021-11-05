package easy_type

func sortArrayByParityII(nums []int) []int {
	res := make([]int, len(nums))
	e := 0
	o := 1
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			res[e] = nums[i]
			e += 2
		} else {
			res[o] = nums[i]
			o += 2
		}
	}
	return res
}
