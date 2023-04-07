package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}
	newNode.next = ll.head
	ll.head = newNode
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
	ll.InsertAtBeginning(5)
	ll.InsertAtBeginning(10)
	ll.InsertAtBeginning(15)
	ll.PrintList()
}
