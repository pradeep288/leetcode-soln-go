package easy_type

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	var capacity int

	// Sort the 2d array by number of units.
	sort.Slice(boxTypes, func(i int, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	for _, box := range boxTypes {
		if truckSize <= box[0] {
			capacity += truckSize * box[1]
			break
		} else {
			capacity +=  box[0] * box[1]
			truckSize -= box[0]
		}
	}

	return capacity
}
