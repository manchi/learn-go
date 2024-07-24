package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("current local time: ", currentTime)

	fmt.Println("starting ticker..")

	ticker := time.Tick(1 * time.Second)
	for i := 0; i < 5; i++ {
		<-ticker // send the ticket to the channel
		fmt.Println("tick:", i)
	}
}
