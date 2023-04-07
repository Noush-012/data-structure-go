package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head  *Node
	limit int
	top   int
}

func (l *LinkedList) push(data int) {
	if l.top == l.limit {
		fmt.Println("stack overflow !!!")
		// return errors.New("stack overflow")
	}
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.top++

}
func (l *LinkedList) pop() {
	if l.top == 0 {
		// return errors.New("stack underflow")
		fmt.Println("stack underflow !!!")
	} else {
		d := l.head.data
		l.head = l.head.next
		l.top--
		fmt.Println("Value:", d, "Deleted Successfully !!!")

	}

}
func (l *LinkedList) peek() {
	if l.top == 0 {
		fmt.Println("No elements to display !!!")
	} else {
		fmt.Println(l.head.data)
	}
}
func (l *LinkedList) printList() {
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func main() {
	l := &LinkedList{}
	l.limit = 5
	l.top = 0
	l.push(10)
	l.push(20)
	l.push(30)
	l.printList()
	// l.push(70)
	// l.push(40)
	// l.push(50)
	l.peek()

}
