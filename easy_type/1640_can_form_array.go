package easy_type

func canFormArray(arr []int, pieces [][]int) bool {
	hashMap := make(map[int][]int)
	for _,val:=range pieces{
		hashMap[val[0]] =val
	}

	for i:=0;i<len(arr);i++{
		if _,ok:=hashMap[arr[i]];ok{
			for _,val:=range hashMap[arr[i]]{
				if arr[i] == val{
					i += 1
				} else{
					return false
				}
			}
		}else{
			return false
		}
	}

	return true
}
