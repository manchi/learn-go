package main

import "fmt"

func main() {
	const MAX = 50
	const THRESHOLD = 20

	for num := 2; num < MAX; num++ {
		isPrime := true

		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Println(num)
		} else {
			continue
		}

		if num > THRESHOLD {
			fmt.Printf("found %v which is greather than %v\n", num, THRESHOLD)
			break
		}
	}
}
