package main

import "fmt"

func main() {
	var num1, num2 float64

	num1 = 44
	num2 = 10

	calc(num1, num2, "+")
	calc(num1, num2, "-")
	calc(num1, num2, "*")
	calc(num1, num2, "/")

	calc(num1, num2, "x")
}

func calc(num1 float64, num2 float64, operator string) {
	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Error: division by zero\n")
		}
	default:
		fmt.Println("Error: invalid operator\n")
	}
}
