package medium_type

import (
	"sort"
)

type location struct {
	id            int
	requiredSeats int
}

func carPooling(trips [][]int, capacity int) bool {
	seatMap := make(map[int]int)
	var l []location

	for i := 0; i < len(trips); i++ {
		seatMap[trips[i][1]] += trips[i][0]
		seatMap[trips[i][2]] -= trips[i][0]
	}

	for k, v := range seatMap {
		l = append(l, location{
			id:            k,
			requiredSeats: v,
		})
	}

	sort.Slice(l, func(i, j int) bool {
		return l[i].id < l[j].id
	})

	var available int
	for _, v := range l {
		available += v.requiredSeats
		if available > capacity {
			return false
		}
	}

	return true
}
