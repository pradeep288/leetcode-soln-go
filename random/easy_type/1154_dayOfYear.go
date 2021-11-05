package easy_type

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	str:= strings.Split(date,"-")
	if len(str)==2{
		res, _:= strconv.Atoi(str[2])
		return res
	}
	return 0
}