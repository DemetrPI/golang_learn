package darts

func Score(x,y float64) int {
	targetCoordinates := x*x + y*y
	switch {
	case targetCoordinates <= 1:
		return 10
	case targetCoordinates <= 25:
		return 5
	case targetCoordinates <= 100:
		return 1
	default:
		return 0
	}
}
