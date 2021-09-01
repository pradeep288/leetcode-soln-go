package medium_type

import (
	"fmt"
	"sort"
)

var subset [][]int

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	bt(0,nums,[]int{})
	fmt.Println(subset)
	return subset
}

func bt(start int, nums []int, cur []int){
	cp := cur
	subset = append(subset, cp)
	for i:=start;i<len(nums);i++{
		if i>start&& nums[i-1] == nums[i]{
			continue
		}
		cur = append(cur, nums[i])
		fmt.Println("Added",cur)
		bt(i+1, nums, cur)
		cur = cur[:len(cur)]
		fmt.Println("Remove",cur)
	}
}
