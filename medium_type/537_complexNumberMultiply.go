package medium_type

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	realA, imgA := parse(a)
	realB, imgB := parse(a)
	realP := realA * realB - imgA * imgB
	img := realA * imgB + realB * imgA

	return strconv.Itoa(realP) + "+" + strconv.Itoa(img) + "i"
}

func parse(s string)(int, int){
	ss := strings.Split(s,"+")
	r,_:=strconv.Atoi(ss[0])
	i,_:=strconv.Atoi(ss[1][:len(ss[1])-1])

	return r, i
}
