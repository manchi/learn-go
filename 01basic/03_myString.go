package main

import "fmt"

func main() {

	// read a number
	var num1 int = 3030
	fmt.Println("your number is: ", num1)

	// read a string
	var name string = "programming"
	fmt.Println("your string input is:", name)
	fmt.Printf("the type of the variable name is: %T\n", name)

	fmt.Printf("this is a float: %.4f\n", 123.123123)

	fmt.Println("a for loop using int x:")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x,", name[i])
	}

	fmt.Println("\na for loop using byte b:")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%b,", name[i])
	}

	fmt.Println("\na for loop using char c:")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c,", name[i])
	}
}
