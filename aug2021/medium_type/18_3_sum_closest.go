package medium_type

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0]+ nums[1] + nums[len(nums)-1]

	for i:=0;i<len(nums)-2;i++{
		lowPointer := i + 1
		highPointer := len(nums)-1

		for lowPointer<highPointer{
			currentSum := nums[i]+ nums[lowPointer] + nums[highPointer]

			if currentSum > target{
				highPointer -= 1
			} else {
				lowPointer += 1
			}
			if currentSum == target{
				return currentSum
			}

			if abs(currentSum-target)< abs(res-target){
				res = currentSum
			}
		}
	}
	return res
}

func abs(val int) int{
	if val < 0{
		return val * -1
	}
	return val
}