package main

import "fmt"

func main() {
	// using arrays
	arr := [3]int{30, 44, 69}
	for i, v := range arr {
		fmt.Printf("array index=%d, value=%d\n", i, v)
	}
	fmt.Println()

	// using slice
	slice := []string{"d", "b", "e"}
	for i, v := range slice {
		fmt.Printf("slice index=%d, value=%s\n", i, v)
	}
	fmt.Println()

	// using string
	str := "golang"
	for i, v := range str {
		fmt.Printf("string index=%d, value=%c\n", i, v)
	}
	fmt.Println()

	// using map
	mymap := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range mymap {
		fmt.Printf("map index=%s, value=%d\n", k, v)
	}
	fmt.Println()

}
