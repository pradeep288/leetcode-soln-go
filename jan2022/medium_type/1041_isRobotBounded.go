package medium_type

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func isRobotBounded(instructions string) bool {
	var curDir Direction
	var curX, curY, count int
	curDir = North
	for count < 4 {
		for _, instruction := range instructions {
			switch string(instruction) {
			case "G":
				switch curDir {
				case North:
					curY += 1
				case South:
					curY -= 1
				case East:
					curX += 1
				case West:
					curX -= 1
				}
			case "L":
				switch curDir {
				case North:
					curDir = West
				case West:
					curDir = South
				case South:
					curDir = East
				case East:
					curDir = North
				}
			case "R":
				switch curDir {
				case North:
					curDir = East
				case East:
					curDir = South
				case South:
					curDir = West
				case West:
					curDir = North
				}
			}
		}
		count += 1
	}
	return curX == curY
}
