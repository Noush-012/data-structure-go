package main

import "fmt"

const max = 5

var stackArr [max]int
var top = -1

func push(data int) {
	if top == max-1 {
		fmt.Println("Stack overflow")
	} else {
		top++
		stackArr[top] = data
	}
}

func pop() {
	if top == -1 {
		fmt.Println(("stack overflow"))
	} else {
		top--
		fmt.Println("Pop successfully!!!")
	}
}

func main() {

}
