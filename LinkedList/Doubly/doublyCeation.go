package Doubly

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (l *DoublyLinkedList) PushBack(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode
}
func (l *DoublyLinkedList) InsertAfter(prevNode *Node, data int) {
    newNode := &Node{data: data}

    newNode.prev = prevNode
    newNode.next = prevNode.next

    if prevNode.next != nil {
        prevNode.next.prev = newNode
    } else {
        l.tail = newNode
    }

    prevNode.next = newNode
}
func main() {
	list := &DoublyLinkedList{}
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(5)
	list.PushBack(6)
	fmt.Println(list)
}