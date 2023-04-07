package main

import (
	"fmt"
	"os"
)

const max = 4

var stackArr [max]int
var top = -1

func push(data int) {
	if top == max-1 {
		fmt.Println("=============       Stack Overflow !!!       =============")
		return
	}
	top++
	stackArr[top] = data
	fmt.Println("\nPushed successfully !!!")
}

func pop() int {
	if top == -1 {
		fmt.Println("\n\n=============       Stack underflow !!!       =============")
		os.Exit(1)
	}
	value := stackArr[top]
	top--
	fmt.Println("Deleted Succefully !!!")
	fmt.Println("Deleted value is:", value)
	return value
}

func print() {
	if top == -1 {
		fmt.Println("\n\n=============       Stack underflow !!!       =============")
		return
	}
	for i := top; i >= 0; i-- {
		fmt.Println("|  ", stackArr[i], "  |")
		if i == 0 {
			fmt.Println("\n********")
		}
	}
}

func main() {
	var choice, data int
	for {
		fmt.Println()
		fmt.Println("1. Push")
		fmt.Println("2. Pop")
		fmt.Println("3. Print the top element")
		fmt.Println("4. Print all the elements of the stack")
		fmt.Println("5. Quit")
		fmt.Print("Please enter your choice:")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			fmt.Print("Enter the element to be pushed:")
			fmt.Scanf("%d", &data)
			push(data)
		case 2:
			pop()
		case 3:
			fmt.Println(stackArr[top])
		case 4:
			print()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("Please enter a valid key !!!")
		}
	}
}
