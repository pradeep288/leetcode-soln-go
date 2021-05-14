// Reference: https://www.youtube.com/watch?v=zseGV4H2EXQ

func checkPossibility(nums []int) bool {
  	max, min := nums[0], nums[len(nums)-1]
	mx , mn := 0, 0

	for i:=0;i<len(nums);i++{
		if nums[i] < max {
			mx += 1
		}
		max = Max(max, nums[i])
	}
	
	
	for i:=len(nums)-1 ;i>=0;i--{
		if nums[i] > min {
			mn += 1
		}
		min = Min(min, nums[i])
	}
	

	if mn <= 1 || mx <= 1{
		return true
	}
	
    return false
}


func Max(x int, y int) int{
	if x > y{
		return x
	}
	return y
}


func Min(x int, y int) int{
	if x < y{
		return x
	}
	return y
}

