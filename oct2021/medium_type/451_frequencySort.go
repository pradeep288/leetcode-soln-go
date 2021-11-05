package medium_type

import (
	"sort"
	"strings"
)

type charInfo struct {
	ch   rune
	freq int
}

func frequencySort(s string) string {
	freq := make(map[rune]int)

	for _, c := range s {
		freq[c]++
	}

	var charInfoList []charInfo

	for k, v := range freq {
		charInfoList = append(charInfoList, charInfo{
			ch:   k,
			freq: v,
		})
	}

	sort.Slice(charInfoList, func(i, j int) bool {
		return charInfoList[i].freq < charInfoList[j].freq
	})

	var res string
	for _, charInfo := range charInfoList {
		res += strings.Repeat(string(charInfo.ch), charInfo.freq)
	}

	return res
}
