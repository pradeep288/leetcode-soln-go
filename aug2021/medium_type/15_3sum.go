package medium_type

import (
	"reflect"
	"sort"
)

var res [][]int
func threeSum(nums []int) [][]int {
	if len(nums) < 3{
		return res
	}
	sort.Ints(nums)
	for i:=0;i<len(nums)-2;i++{
		l:= i+1
		h:=len(nums)-1
		for l<h{
			currSum:= nums[i] + nums[l] + nums[h]

			if currSum == 0{
				newTriplet:=[]int{nums[i], nums[l], nums[h]}
				if  !tripleExists(newTriplet) {
					res = append(res, newTriplet)
				}
				h -= 1
				l += 1
			}
			if currSum > 0 {
				h -= 1
			}
			if currSum < 0{
				l += 1
			}
		}
	}
	return res
}

func tripleExists(triplet []int) bool{
	for _, item:=range res{
		if reflect.DeepEqual(item, triplet){
			return true
		}
	}
	return false
}