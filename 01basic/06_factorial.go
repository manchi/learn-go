package main

import (
	"errors"
	"fmt"
)

func main() {
	result1, err := factorial(6)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("factorial of %d is %d\n", 6, result1)

	result2, err := factorial(-6)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("factorial of %d is %d\n", -6, result2)
	}
}

func factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is undefined for negative number")
	}

	if n == 0 {
		return 1, nil
	}
	fact, err := factorial(n - 1)
	if err != nil {
		return 0, err
	}
	return n * fact, nil
}
