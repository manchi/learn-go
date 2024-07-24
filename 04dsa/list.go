package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("using container/list")

	l := list.New()

	l.PushFront(10)
	l.PushFront(20)
	l.PushBack(30)
	l.PushBack(40)
	printList(l)

	var el *list.Element // find element
	for el = l.Front(); el.Value != 30; el = el.Next() {
	}

	l.InsertAfter(44, el)  // insert after 30
	l.InsertBefore(22, el) // insert before 30
	fmt.Println("total size: ", l.Len())
	printList(l)
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v, ", e.Value)
	}
	fmt.Println()
}
