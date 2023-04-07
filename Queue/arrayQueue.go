package main

import "fmt"

const max = 5

var queue [max]int
var front, rear int = -1, -1

func Enqueue(data int) {
	if rear == max-1 {
		fmt.Println("Queue overflow !!!")
	} else if front == -1 && rear == -1 {
		front = 0
		rear = 0
		queue[rear] = data
	} else {
		rear++
		queue[rear] = data
	}
}
func Dequeue() {
	if front == -1 && rear == -1 {
		fmt.Println("Queue underflow !!!")
	} else if front == rear {
		front = -1
		rear = -1
	} else {
		d := queue[front]
		fmt.Println("Value:", d, "Dequeued Successfully !!!")
		front++
	}
}
func peek() {
	if front == -1 && rear == -1 {
		fmt.Println("Queue underflow !!!")
	} else {
		fmt.Println(queue[front])
	}
}
func isEmpty() {
	if front == -1 && rear == -1 {
		fmt.Println("Queue is empty !!!")
	} else {
		fmt.Println("Queue is not empty !!!")
	}
}
func printQ() {
	if front == -1 && rear == -1 {
		fmt.Println("Queue is empty !!!")
	} else {
		for i := front; i < rear; i++ {
			fmt.Print(queue[i], " _ ")
		}
	}
}

func main() {
	Enqueue(10)
	Enqueue(-1)
	Enqueue(4)
	Enqueue(-5)

	// Dequeue()
	printQ()
	isEmpty()
	peek()

}
