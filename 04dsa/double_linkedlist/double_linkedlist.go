package double_linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	//tail   *Node
	length int
}

func (l *LinkedList) InsertAtHead(data int) {
	newNode := &Node{data: data, next: nil, prev: nil}

	if l.head == nil {
		l.head = newNode
	} else {
		temp := l.head

		newNode.next = temp
		temp.prev = newNode

		l.head = newNode
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
		newNode.prev = current
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
		newNode.prev = current

		newNode.next = temp
		temp.prev = newNode

		// inc count
		l.length++
	}
}

func (l *LinkedList) DeleteAtHead() {
	head := l.head

	temp := head.next
	temp.prev = nil

	l.head = temp

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
		temp.next.prev = temp.prev

		l.length--
	}
}

func (l *LinkedList) Reverse() {
	var curr, prev, next *Node
	curr = l.head
	prev = nil

	for curr != nil {
		prev = curr.prev
		next = curr.next

		curr.next = prev
		curr.prev = next

		prev = curr
		curr = next
	}
	l.head = prev
}

func Print(l *LinkedList) {

	head := l.head
	for head != nil {
		fmt.Printf("%v -> ", head.data)
		head = head.next
	}
	fmt.Printf("nil\n")
}
