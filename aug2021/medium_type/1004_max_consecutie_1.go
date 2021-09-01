package medium_type

import "fmt"

func main(){
	nums:=[]int {1,1,1,0,0,0,1,1,1,1,0}
	fmt.Println(longestOnes(nums,2))
}
func longestOnes(nums []int, k int) int {
	var fp, sp int

	for fp=0; fp< len(nums);fp++{
		if nums[fp] == 0{
			k -= 1
		}
		if k<0{
			if nums[sp] == 0 {
				k+=1
			}
			sp+=1
		}
	}
	return fp-sp
}

/*
1 , 1, 1, 0, 0, 0, 1, 1, 1, 1, 0
               fp
  sp

k=-1
 */