package main

import (
	"fmt"
)

const MAX = 5

var stack [MAX]int
var top = -1

func push(data int) {

	if top == -1 {

		fmt.Println(" Stack underflow !!!")
	} else {

		top++
		stack[top] = data
		fmt.Println(" Pushed successfully!!!")
	}
}
func main() {

	var arr []int
	var sum int

	arr = []int{1, 2, 3, 4, 6, 8}

	for i := 0; i < len(arr); i++ {

		sum = arr[i] + sum
	}
	fmt.Println(sum)
}
