package medium_type

import "strings"

type MapSum struct {
	m map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		m: make(map[string]int),
	}
}


func (this *MapSum) Insert(key string, val int)  {
	if _,ok:=this.m[key];ok{
		this.m[key] = val
	} else {
		this.m[key] = val
	}
}


func (this *MapSum) Sum(prefix string) int {
	var sum int
	for k,v:=range this.m {
		if strings.HasPrefix(k, prefix){
			sum += v
		}
	}
	return sum
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
