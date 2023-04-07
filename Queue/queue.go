package main

import "fmt"

// Node is a structure to hold the value and pointer to next node.
type Node struct {
	Value int
	Next  *Node
}

// Queue is a structure to hold the pointer to first and last node.
type Queue struct {
	First *Node
	Last  *Node
}

// Enqueue adds a new node to the end of the queue.
func (q *Queue) Enqueue(value int) {
	newNode := &Node{value, nil}
	if q.Last == nil {
		q.Last = newNode
		q.First = newNode
	} else {
		q.Last.Next = newNode
		q.Last = newNode
	}
}

// Dequeue removes the first node from the queue and returns its value.
func (q *Queue) Dequeue() int {
	if q.First == nil {
		fmt.Println("Queue is empty.")
		return -1
	}
	value := q.First.Value
	q.First = q.First.Next
	if q.First == nil {
		q.Last = nil
	}
	return value
}

func main() {
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
