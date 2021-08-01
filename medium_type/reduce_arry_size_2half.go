package medium_type

import "sort"

func minSetSize(arr []int) int {
	freq:= map[int]int{}
	for i:=0;i<len(arr);i++{
		if _,ok:= freq[arr[i]];ok{
			freq[arr[i]] += 1
		} else {
			freq[arr[i]] = 1
		}
	}

	var val []int
	for _, v:=range freq{
		val=append(val, v)
	}
	sort.Ints(val)

	rem, cur, res:=len(arr),len(arr), 0

	for i:=len(val)-1;i>=0;i--{
		if cur <= rem/2{
			break
		} else {
			cur -= val[i]
			res += 1
		}
	}
	return res
}