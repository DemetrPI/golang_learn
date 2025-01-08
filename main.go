package main

import (
	"fmt"
	"golang_learn/greetings"
	"golang_learn/lasagna"

)

func main() {
	fmt.Println(greetings.HelloWorld())

	// Use the lasagna package
    actualMinutesInOven := 30
    // numberOfLayers := 3

    fmt.Printf("Remaining Oven Time: %d minutes\n", lasagna.RemainingOvenTime(actualMinutesInOven))
    // fmt.Printf("Preparation Time: %d minutes\n", lasagna.PreparationTime(numberOfLayers))
    // fmt.Printf("Elapsed Time: %d minutes\n", lasagna.ElapsedTime(numberOfLayers, actualMinutesInOven))
}

