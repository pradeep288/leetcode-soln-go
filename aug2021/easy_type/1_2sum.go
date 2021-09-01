package easy_type


func twoSum(nums []int, target int) [2]int {
	var res [2]int
	dict:=make(map[int]int)
	for i:=0;i<len(nums)-1;i++{
		sum:= target-nums[i]

		if _,ok:=dict[sum];ok{
			res[0] = i
			res[1] = dict[sum]
			return res
		} else {
			dict[nums[i]] = i
		}
	}


	return res
}