package medium_type

import "sort"

func groupAnagrams(strs []string) [][]string {
	var res [][]string

	if len(strs)==0{
		return res
	}

	hashMap:= make(map[string][]string)

	for i:=0;i<len(strs);i++{
		curStr :=strs[i]
		//convert string to rune
		r:= []rune(strs[i])
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		s := string(r)
		hashMap[s] =append(hashMap[s], curStr)
	}

	for _,v:= range hashMap{
		res = append(res, v)
	}

	return res
}