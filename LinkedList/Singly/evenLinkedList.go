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
		if current.value%2 == 0 {
			fmt.Print(" ~ ", current.value)
		}
		current = current.next
	}
}
func main() {
	list := &LinkedList{}
	for i := 1; i <= 10; i++ {
		list.Append(i)
	}

	list.Display()
}
