package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	RemainingOvenTime := OvenTime - actualMinutesInOven
	return RemainingOvenTime
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
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

// // ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
// func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
// 	ElapsedTime := PreparationTime(numberOfLayers) + actualMinutesInOven

// 	return ElapsedTime

// }
