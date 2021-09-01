package medium_type

import "strings"

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	vacant:=1
	for _,node:= range nodes{
		vacant -= 1
		if node!="#"{
			vacant += 2
		}

	}
	return vacant == 0
}
