package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePrepTime int) int {
	if averagePrepTime == 0 {
		averagePrepTime = 2
	}

	return len(layers) * averagePrepTime
}

// TODO: define the 'Quantities()' function
// Using naked return
func Quantities(layers []string) (noodleQuantity int, sauceQuantity float64) {
	noodleQuantity = 0
	sauceQuantity = 0.0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodleQuantity += 50
		}
		if layers[i] == "sauce" {
			sauceQuantity += 0.2
		}

	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var adjustedQuant []float64

	for i := 0; i < len(quantities); i++ {
		adjustedQuant = append(adjustedQuant, (quantities[i]/2)*float64(portions))
	}

	return adjustedQuant
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
