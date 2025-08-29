package space


type Planet string

func Age(seconds float64, planet Planet) (float64) {
	var earthYearSeconds float64 = 31557600
	standard :=  seconds /earthYearSeconds
	switch planet {
	case "Mercury":
		return standard / 0.2408467
	case "Venus":
		return standard / 0.61519726
	case "Earth":
		return standard
	case "Mars":
		return standard / 1.8808158
	case "Jupiter":
		return standard / 11.862615
	case "Saturn":
		return standard / 29.447498
	case "Uranus":
		return standard / 84.016846
	case "Neptune":
		return standard / 164.79132
	default:
		return -1
	}

}
