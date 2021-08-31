package easy_type

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	var res string
	lenN1, lenN2 := len(num1)-1, len(num2)-1
	i, j := lenN1, lenN2
	var carry, div uint8
	carry = 0
	div = 10

	for i>=0 || j>=0{
		sum := carry
		if i>=0{
			sum += num1[i] - 48
			i-=1
		}
		if j>=0{
			sum += num2[j] - '0'
			j-=1
		}
		carry = sum/10
		res = ("0"+ string(+ sum%div)) + res
	}

	if carry>0{
		res = string(carry) + res
	}
	fmt.Println(res)

	return res
}
