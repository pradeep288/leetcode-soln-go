package medium_type

import (
	"math"
	"sort"
)

// Point represents Point in x-y plane
type Point struct {
	distanceFromOrigin float64
	xCoordinate        int
	yCoordinate        int
}

func kClosest(points [][]int, k int) [][]int {
	var res [][]int
	var p []Point

	// Represent the given input using Point struct.
	for _, v := range points {
		d := math.Sqrt(math.Pow(float64(v[0])-float64(0), float64(2)) +
			math.Pow(float64(v[1])-float64(0), float64(2)))
		p = append(p, Point{
			distanceFromOrigin: d,
			xCoordinate:        v[0],
			yCoordinate:        v[1],
		})
	}

	// Sort the Point struct based on the distance from origin
	sort.Slice(p, func(i, j int) bool {
		return p[i].distanceFromOrigin < p[j].distanceFromOrigin
	})

	// Return first k Points closest to the origin
	for i := 0; i < k; i++ {
		res = append(res, []int{p[i].xCoordinate, p[i].yCoordinate})
	}

	return res
}
