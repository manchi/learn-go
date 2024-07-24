package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func main() {
	start := time.Now()
	id := getUserByName1("hari")
	fmt.Println("id:", id)

	wg := &sync.WaitGroup{}

	ch := make(chan *Message, 2)
	wg.Add(2)

	go getUserChats1(id, ch, wg)
	//log.Println(chats)

	go getUserFriends1(id, ch, wg)
	//log.Println(friends)

	wg.Wait()
	close(ch)

	for msg := range ch {
		log.Println(msg)
	}
	log.Println("total time:", time.Since(start))
}

func getUserFriends1(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 4)

	ch <- &Message{
		friends: []string{
			"sagar",
			"bharath",
			"shankar",
		},
	}
	wg.Done()
}

func getUserChats1(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 3)

	ch <- &Message{
		chats: []string{
			"ram",
			"krishna",
			"ravi",
		}}
	wg.Done()
}

func getUserByName1(name string) string {
	time.Sleep(time.Second * 2)

	return "s-010"
}
