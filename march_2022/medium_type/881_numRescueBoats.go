package medium_type

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	var boatCount int
	i, j := 0, len(people)-1
	for i <= j {
		if people[i]+people[j] > limit {
			j -= 1
		} else {
			i += 1
			j -= 1
		}

		boatCount += 1
	}
	return boatCount
}
