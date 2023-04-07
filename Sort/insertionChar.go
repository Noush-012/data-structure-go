package main

import "fmt"

func insertionSort(char []rune) {
	n := len(char)
	for i := 1; i < n; i++ {
		key := char[i]
		j := i - 1
		for j >= 0 && char[j] > key {
			char[j+1] = char[j]
			j = j - 1
		}
		char[j+1] = key
	}
}

func main() {
	char := []rune{'d', 'a', 'c', 'b'}
	fmt.Println("Before sorting: ", string(char))
	insertionSort(char)
	fmt.Println("After sorting: ", string(char))
}
