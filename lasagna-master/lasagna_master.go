package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layer []string, averagePrepTime int) int {

	if averagePrepTime == 0 {
		averagePrepTime = 2
	}

	return len(layer) * averagePrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layer []string) (noodles int, sauce float64) {
	for v := range layer {
		if layer[v] == "sauce" {
			sauce += 0.2
		} else if layer[v] == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(friendsList)] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(slice []float64, noOfPortions int) []float64 {
	noOfPortion := float64(noOfPortions) / 2
	newSlice := make([]float64, 0, len(slice))
	for i := range slice {
		newSlice = append(newSlice, slice[i]*float64(noOfPortion))
	}

	return newSlice
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
