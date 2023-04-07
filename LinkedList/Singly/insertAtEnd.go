package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if ll.tail == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		ll.tail = newNode
	}
}

func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func main() {
	ll := &LinkedList{}
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(4)
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(6)
	ll.InsertAtEnd(5)
	ll.PrintList()
}
