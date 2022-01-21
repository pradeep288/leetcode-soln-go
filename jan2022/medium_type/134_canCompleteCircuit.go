package medium_type

func canCompleteCircuit(gas []int, cost []int) int {
	var sum func([]int) int

	sum = func(s []int) int {
		var total int
		for i := 0; i < len(s); i++ {
			total += s[i]
		}
		return total
	}

	if sum(gas) < sum(cost) {
		return -1
	}

	var startPoint, total int
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		if total < 0 {
			startPoint = i + 1
			total = 0
		}
	}
	return startPoint
}
