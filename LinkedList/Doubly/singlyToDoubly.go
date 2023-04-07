package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}
	newNode.next = ll.head
	ll.head = newNode
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

// convert the singly linked list to doubly linked list
func (ll *LinkedList) convertToDoublyLinkedList() {
	current := ll.head
	for current != nil {
		current.prev = current.next
		current = current.next
	}
}
func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Println(current.data)
		current = current.prev
	}
}

func main() {
	ll := &LinkedList{}
	ll.InsertAtBeginning(5)
	ll.InsertAtBeginning(10)
	ll.InsertAtEnd(15)
	ll.convertToDoublyLinkedList()
	ll.PrintList()

}
