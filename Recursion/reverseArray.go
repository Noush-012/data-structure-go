package main

import "fmt"

func reverseArray(arr []int, index int) {
	if index == len(arr) {
		return
	}
	reverseArray(arr, index+1)
	fmt.Print(arr[index], " ")
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	reverseArray(arr, 0) 
	
}
