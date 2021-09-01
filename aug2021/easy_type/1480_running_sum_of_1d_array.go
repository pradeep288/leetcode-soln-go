func runningSum(nums []int) []int {
    sum := []int{}
	sum = append(sum, nums[0])
	for i:=1;i<len(nums);i++{
		val := sum[i-1] + nums[i]
		sum = append(sum, val)
	}
    
    return sum
}
