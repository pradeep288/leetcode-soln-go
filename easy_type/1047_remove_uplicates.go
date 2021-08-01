package easy_type

import "strings"

func removeDuplicates(s string) string {
	res:= []string{}
	for _,v:= range s{
		if len(res) > 0 && res[len(res)-1] == string(v){
			res = res[:len(res)-1]
		} else  {
			res =append(res, string(v))
		}
	}
	return strings.Join(res,"")
}
