package main

import (
	"fmt"
)

func main() {
	fmt.Println("map test")

	map1 := make(map[string]int)
	map1["sunday"] = 1
	map1["monday"] = 2
	map1["tuesday"] = 3
	map1["wednesday"] = 4
	map1["thursday"] = 5
	map1["friday"] = 6
	map1["saturday"] = 7

	fmt.Println(map1)

	map1["weekday"] = 7
	fmt.Println(map1)

	map1["holiday"] = 7
	fmt.Println(map1)
}
