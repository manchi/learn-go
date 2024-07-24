package main

import (
	"fmt"
)

func main() {
	var a = 10394
	fmt.Printf("%d\n", a)
	fmt.Printf("the type of the variable is: %T\n", a)

	var b = 3.14159
	fmt.Printf("%f\n", b)
	fmt.Printf("the type of the variable is: %T\n", b)

	var c = false
	fmt.Printf("%t\n", c)
	fmt.Printf("the type of the variable is: %T\n", c)

	var d = "test string"
	fmt.Println(d)
	fmt.Printf("the type of the variable is: %T\n", d)

}
