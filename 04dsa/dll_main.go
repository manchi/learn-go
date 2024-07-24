package main

import (
	"dsa/double_linkedlist"
	"fmt"
)

func main() {

	fmt.Println("double_linkedlist.LinkedList{}")
	list := &double_linkedlist.LinkedList{}

	list.InsertAtHead(20)
	list.InsertAtHead(10)
	list.InsertAtHead(34)
	double_linkedlist.Print(list)

	list.InsertAtTail(80)
	list.InsertAtTail(90)
	list.InsertAtTail(15)
	double_linkedlist.Print(list)

	list.Insert(2, 44)
	double_linkedlist.Print(list)

	list.Insert(3, 66)
	double_linkedlist.Print(list)

	list.Insert(4, 70)
	double_linkedlist.Print(list)

	list.Delete(5)
	double_linkedlist.Print(list)

	list.DeleteAtTail()
	list.DeleteAtTail()
	double_linkedlist.Print(list)

	list.DeleteAtHead()
	list.DeleteAtHead()
	double_linkedlist.Print(list)

	list.Reverse()
	double_linkedlist.Print(list)
}
