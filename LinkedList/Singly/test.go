package main

import "fmt"

// Node represents a node in the linked list
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	Head *Node
}

// Insert adds a new node to the linked list
func (l *LinkedList) Insert(value int) {
	node := &Node{Value: value, Next: nil}
	if l.Head == nil {
		l.Head = node
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func main() {
	linkedList := &LinkedList{}
	linkedList.Insert(1)
	linkedList.Insert(2)
	linkedList.Insert(3)
	linkedList.Insert(4)


	node := linkedList.Head
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}