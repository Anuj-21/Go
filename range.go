package main

import "fmt"

func main() {
	//create a slice
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	//print the numbers
	for i := range numbers { //for array range return the index of the item as integer.
		fmt.Println("Slice item", i, "is", numbers[i])
	}

	//create a map
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	//print map using keys
	for country := range countryCapitalMap { //for maps, range returns the key of the next key-value pair
		fmt.Println("capital of ", country, "is", countryCapitalMap[country])
	}

	//print map using key-value

}
