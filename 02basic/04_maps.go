package main

import "fmt"

func main() {
	var fruits = make(map[string]string)
	fruits["apple"] = "red"
	fruits["banana"] = "yellow"
	fruits["grapes"] = "purple"

	fmt.Println(fruits)
	for fruit, color := range fruits {
		fmt.Printf("fruit %v has %v color\n", fruit, color)
	}

	var cities = make(map[string]int)
	cities["NY"] = 11
	cities["CA"] = 22
	cities["MN"] = 33

	fmt.Println(cities)

	for city, rank := range cities {
		fmt.Printf("city %v has rank - %v\n", city, rank)
	}

	// update
	fruits["banana"] = "red"
	fmt.Println(fruits)

	// delete
	delete(fruits, "banana")
	fmt.Println(fruits)
}
