package main

import (
	"dsa/linkedlist"
	"fmt"
)

func main() {
	fmt.Println("linkedlist.LinkedList{}")
	list := &linkedlist.LinkedList{}

	list.InsertAtHead(90)
	list.InsertAtHead(13)
	list.InsertAtHead(34)
	linkedlist.Print(list)

	list.InsertAtTail(22)
	list.InsertAtTail(89)
	list.InsertAtTail(15)
	linkedlist.Print(list)

	list.Insert(2, 44)
	linkedlist.Print(list)

	list.Insert(3, 66)
	linkedlist.Print(list)

	list.Insert(4, 99)
	linkedlist.Print(list)

	list.Delete(5)
	linkedlist.Print(list)

	list.DeleteAtTail()
	list.DeleteAtTail()
	linkedlist.Print(list)

	list.DeleteAtHead()
	list.DeleteAtHead()
	linkedlist.Print(list)

	list.Reverse()
	linkedlist.Print(list)
}
