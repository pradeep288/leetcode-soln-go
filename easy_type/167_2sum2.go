package easy_type

func twoSum2(nums []int, target int) []int {
	var res []int


	low,high :=0,len(nums)-1

	for low<high{
		sum := nums[low]+nums[high]
		if  sum== target{
			return []int{low+1,high+1}
		}

		if sum > target{
			high -=1
		}

		if sum < target{
			low +=1
		}
	}
	return res
}