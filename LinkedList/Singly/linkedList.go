package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	l.size++
}
func (l *LinkedList) Display() {
	current := l.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
func main() {
	list := &LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(10)
	// fmt.Println(list)
	list.Display()
}
