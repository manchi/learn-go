package main

import "fmt"

func main() {
	arr1 := [6]int{20, 34, 54, 77, 22, 85}

	// using index and len()
	sum1 := 0
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("{%d:%d} ", i, arr1[i])
		sum1 += arr1[i]
	}
	fmt.Printf("\nsum1: %v\n", sum1)

	// using range
	arr2 := [6]int{10, 14, 5, 92, 56, 89}
	sum2 := 0
	for i, v := range arr2 {
		fmt.Printf("{%d:%d} ", i, v)
		sum2 += v
	}
	fmt.Printf("\nsum2: %v\n", sum2)

	arr3 := [6]int{57, 2, 5, 66, 93, 23}

	// pass by value
	arr4 := arr3

	arr4[3] = 55               // This change won't affect arr3
	fmt.Println("arr3:", arr3) //[57 2 5 66 93 23]
	fmt.Println("arr4:", arr4) //[57 2 5 55 93 23]
}
