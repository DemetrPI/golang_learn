package lasagna


func PreparationTime(layers []string, numberOfLayers int) int {
	if numberOfLayers == 0 {
		numberOfLayers = 2
	}
	PreparationTime := len(layers) * numberOfLayers
	return PreparationTime
}

func Quantities(layers []string) (int, float64) {
	var sauce float64
	var noodles int

	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			sauce += 0.2
		} else if layers[i] == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(quantities []float64, scale int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i, quantity := range quantities {
		scaledQuantities[i] = quantity / 2 * float64(scale)
	}
	return scaledQuantities

}