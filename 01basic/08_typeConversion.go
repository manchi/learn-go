package main

import (
	"fmt"
	"strconv"
)

func convertToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	} else {
		return num, err
	}
}

func convertToFloat(s string) (float64, error) {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	} else {
		return num, err
	}
}

func main() {
	strInt := "10394"
	strFloat := "3.14159"

	intNum, err := convertToInt(strInt)
	if err != nil {
		fmt.Println("error converting int", err)
	} else {
		fmt.Println("convert to int: ")
		fmt.Printf("%d\n", intNum)
		fmt.Printf("the type of the variable is: %T\n", intNum)
	}

	floatNum, err := convertToFloat(strFloat)
	if err != nil {
		fmt.Println("error converting float", err)
	} else {
		fmt.Println("convert to float: ")
		fmt.Printf("%f\n", floatNum)
		fmt.Printf("the type of the variable is: %T\n", floatNum)
	}
}
