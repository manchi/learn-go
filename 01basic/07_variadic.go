package main

import "fmt"

func main() {

	fmt.Println(average(1, 2, 3, 4, 5))
	fmt.Println(average(14, 62, 23, 34, 15))
	fmt.Println(average(1.2, 2.1, 3.4, 4.3, 5.6))
	fmt.Println(average())
}

func average(nums ...float64) float64 {
	total := 0.0
	count := 0
	for _, num := range nums {
		total += num
		count++
	}

	if count == 0 {
		return 0
	}
	return total / float64(count)
}
