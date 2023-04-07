package main

import "fmt"

func main() {

	Q := &Queue{}
	Q.Enqueue(10)
	Q.Dequeue()
}

type Queue struct {
	enqStack []int
	deqStack []int
}

func (q *Queue) Enqueue(val int) {
	q.enqStack = append(q.enqStack, val)
	fmt.Println("Success!")
}

func (q *Queue) Dequeue() int {
	if len(q.deqStack) == 0 {
		for len(q.enqStack) > 0 {
			// Pop from the enqStack and push onto the deqStack
			n := len(q.enqStack) - 1
			q.deqStack = append(q.deqStack, q.enqStack[n])
			q.enqStack = q.enqStack[:n]
			fmt.Println("Success!")
		}
	}

	// Pop from the deqStack to dequeue an element
	n := len(q.deqStack) - 1
	if n < 0 {
		return -1 // Queue is empty
	}
	val := q.deqStack[n]
	q.deqStack = q.deqStack[:n]
	return val
}
