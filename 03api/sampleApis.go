package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	id := getUserByName("hari")
	fmt.Println("id:", id)

	chats := getUserChats(id)
	log.Println(chats)

	friends := getUserFriends(id)
	log.Println(friends)

	log.Println("total time:", time.Since(start))
}

func getUserFriends(id string) []string {
	time.Sleep(time.Second * 4)

	return []string{
		"sagar",
		"bharath",
		"shankar",
	}

}

func getUserChats(id string) []string {
	time.Sleep(time.Second * 3)

	return []string{
		"ram",
		"krishna",
		"ravi",
	}
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 2)

	return "s-010"
}
