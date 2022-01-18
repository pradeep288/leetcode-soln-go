package easy_type

func canPlaceFlowers(flowerbed []int, n int) bool {
	var curPlot int

	// Special Case1:  when number of flowers to plant is 0
	if n == 0 {
		return true
	}
	// Special Case2:  when length of flowerbed is 1
	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 {
			n--
		}
		return n == 0
	}

	for curPlot < len(flowerbed) && n != 0 {
		switch curPlot {
		// Special Case3:  When Plot is at left Edge
		case 0:
			if flowerbed[curPlot] == 0 && flowerbed[curPlot+1] == 0 {
				flowerbed[curPlot] = 1
				n--
			}
		// Special Case4:  When Plot is at right Edge
		case len(flowerbed) - 1:
			if flowerbed[curPlot] == 0 && flowerbed[curPlot-1] == 0 {
				flowerbed[curPlot] = 1
				n--
			}
		default:
			if flowerbed[curPlot] == 0 && flowerbed[curPlot+1] == 0 && flowerbed[curPlot-1] == 0 {
				flowerbed[curPlot] = 1
				n--
			}
		}
		curPlot++
	}

	return n == 0
}
