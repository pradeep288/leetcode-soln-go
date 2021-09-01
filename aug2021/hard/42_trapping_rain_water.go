package hard

func trap(height []int) int {
	size:=len(height)

	leftMax, rightMax := make([]int,size),make([]int,size)

	var water int

	for i:=0;i<size;i++{
		if i==0{
			leftMax[i] =height[i]
			rightMax[size-1] = height[size-1]
		} else {
			leftMax[i] = maxE(height[i], leftMax[i-1])
			rightMax[size-i] = maxE(height[size-i-1], rightMax[size-1])
		}
	}

	for i:=0;i<size;i++{
		water += minE(rightMax[i],leftMax[i]) - height[i]
	}
	return water
}

func minE(x int, y int) int {
	if x<y {
		return x
	}
	return y
}

func maxE(x int, y int) int {
	if x > y {
		return x
	}
	return y
}