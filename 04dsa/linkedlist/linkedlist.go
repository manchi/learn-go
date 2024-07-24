package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) InsertAtHead(data int) {
	newNode := &Node{data: data, next: nil}

	if l.head == nil {
		l.head = newNode
	} else {
		temp := l.head
		l.head = newNode
		newNode.next = temp
	}
	l.length++
}

func (l *LinkedList) InsertAtTail(data int) {
	newNode := &Node{data: data, next: nil}

	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		// go to last node
		for current.next != nil {
			current = current.next
		}
		// set the next node as new node
		current.next = newNode
	}
	l.length++
}

func (l *LinkedList) Insert(n int, data int) {
	newNode := &Node{data: data, next: nil}

	if n == 0 {
		l.InsertAtHead(data)
	} else if n == l.length-1 {
		l.InsertAtTail(data)
	} else {
		// go to n th current
		current := l.head
		for i := 0; i < n-1; i++ {
			current = current.next
		}
		temp := current.next
		current.next = newNode
		newNode.next = temp

		// inc count
		l.length++
	}
}

func (l *LinkedList) DeleteAtHead() {
	head := l.head
	l.head = head.next

	l.length--
}

func (l *LinkedList) DeleteAtTail() {
	var temp *Node = nil
	head := l.head

	// go to last node
	for head.next != nil {
		temp = head
		head = head.next
	}
	// set the next node as new node
	temp.next = nil

	l.length--
}

func (l *LinkedList) Delete(n int) {

	if n == 0 {
		l.DeleteAtHead()
	} else if n == l.length-1 {
		l.DeleteAtTail()
	} else {
		current := l.head
		// go to the n-1 node
		for i := 0; i < n-1; i++ {
			current = current.next
		}
		temp := current.next
		current.next = temp.next

		l.length--
	}
}

func (l *LinkedList) Reverse() {
	/*
		Method
		A method is a function that is associated with a specific type. It has a receiver argument, which is the
		instance of the type on which the method is called. This is similar to instance methods in Java.

		When to Use Methods

		- Object-Oriented Style: Use methods when you want to attach behavior to a specific type. This is akin to
		defining instance methods in an object-oriented programming language like Java.
		- Encapsulation: Methods are useful for encapsulating behavior that logically belongs to a type. For example,
		methods that operate on the data contained within a struct.
		- Interfaces: Methods are essential for defining and implementing interfaces in Go. A type must implement all
		methods of an interface to be considered as implementing that interface.
	*/
	var curr, prev, next *Node
	curr = l.head
	prev = nil

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

// This is a Standalone Function
func Print(l *LinkedList) {
	/*
		Standalone Function

		A standalone function is not associated with any particular type. It is defined at the package level and can be
		called from anywhere within the package (or outside if it is exported).

		When to Use Standalone Functions
		- Utility Functions: Use standalone functions for generic operations that do not naturally belong to a specific type.
		Examples include mathematical operations, string manipulations, or other utility functions.
		- Functional Programming Style: When following a functional programming paradigm, where functions are treated as
		first-class citizens and passed around as arguments, standalone functions are preferred.
	*/
	head := l.head
	for head != nil {
		fmt.Printf("%v -> ", head.data)
		head = head.next
	}
	fmt.Printf("nil\n")
}

func main() {

	list := &LinkedList{}

	list.InsertAtHead(90)
	list.InsertAtHead(13)
	list.InsertAtHead(34)
	Print(list)

	list.InsertAtTail(22)
	list.InsertAtTail(89)
	list.InsertAtTail(15)
	Print(list)

	list.Insert(2, 44)
	Print(list)

	list.Insert(3, 66)
	Print(list)

	list.Insert(4, 99)
	Print(list)

	list.Delete(5)
	Print(list)

	list.DeleteAtTail()
	list.DeleteAtTail()
	Print(list)

	list.DeleteAtHead()
	list.DeleteAtHead()
	Print(list)

	list.Reverse()
	Print(list)
}
