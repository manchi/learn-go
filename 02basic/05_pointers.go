package main

import "fmt"

func modifyValue(val *int, newVal int) {
	*val = newVal
}
func main() {
	fmt.Println("pointers")

	x := 10
	fmt.Println("original value is", x)

	ptr := &x
	fmt.Println("memory address is", ptr)

	fmt.Println("value at address is", *ptr)

	// modify
	*ptr = 100
	fmt.Println("updated value at address is", *ptr)

	modifyValue(ptr, 200)
	fmt.Println("updated value at address is", *ptr)
}
