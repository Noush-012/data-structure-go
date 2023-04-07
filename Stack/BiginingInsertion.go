package main

import (
	"fmt"
)

const MAX = 5

var stackArr [MAX]int
var first int

func push(data int) {
	if first == MAX-1 {
		fmt.Println("=============       Stack Overflow !!!       =============")
		return
	}
	first++
	for i := first; i > 0; i-- {
		stackArr[i] = stackArr[i-1]
	}
	stackArr[0] = data
	fmt.Println("\nPushed successfully !!!")
}

func print() {
	if first == -1 {
		fmt.Println("\n\n=============       Stack underflow !!!       =============")
		return
	}
	for i := 0; i <= first; i++ {
		fmt.Println("|  ", stackArr[i], "  |")
		if i == 0 {
			fmt.Println("\n********")
		}
	}
}
