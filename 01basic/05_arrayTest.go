package main

import (
	"log"
)

func main() {
	myArray := [4]int{1, 2, 5, 5}
	length := len(myArray)
	log.Println("total length:", length)

	log.Println("content of 1D array")
	for i := 0; i < length; i++ {
		log.Printf("%d, ", myArray[i])
	}
	log.Println()
	log.Println("__________")

	// another array - dynamic length
	array2 := [...]int{2, 15, 26, 34, 55, 67}
	log.Println("content of 1D array (dynamic)")
	for _, num := range array2 {
		log.Printf("%d, ", num)
	}
	log.Println()
	log.Println("__________")

	// print 2D array
	twoD := [3][3]int{
		{1, 3, 5},
		{6, 8, 9},
		{12, 33, 55},
	}
	log.Println("content of 2D array")
	for _, num := range twoD {
		for _, other := range num {
			log.Printf("%d, ", other)
		}
		log.Println("..")
	}
	log.Println("__________")
	log.Println()

	// using slice
	slice1 := []int{44, 2, 45, 76, 87}
	log.Println("using slice")
	log.Println(slice1)

	log.Println("__________")
	// using make func
	slice2 := make([]int, 4)
	log.Println("using make func")
	log.Println(slice2)
	log.Println("slice append 100")
	slice2 = append(slice2, 100)
	log.Println(slice2)

	log.Println("__________")
	// capacity
	log.Println("cap built-in function (the capacity of v, according to its type) is:", cap(slice2))

	log.Println("__________")
	// len
	log.Println("returns the length of v, according to its type is:", len(slice2))
}
