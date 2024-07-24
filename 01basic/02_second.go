package main

import "fmt"

func main() {
	fmt.Println("Hello World from Golang !!")

	// read a number
	fmt.Println("give me a number:")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println("your number is: ", num1)

	// read a string
	fmt.Println("give me a string:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("your string input is:", name)

}
