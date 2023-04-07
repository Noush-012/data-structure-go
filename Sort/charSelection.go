package main

import (
	"fmt"
)

func selectionSort(char []rune) {
	n := len(char)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; i < n; j++ {
			if char[j] < char[min] {
				min = j

			}
		}
		char[i], char[min] = char[min], char[i]
	}
}
func main() {

	char := []rune{'d', 'c', 'a', 'b'}
	selectionSort(char)
	fmt.Println(string(char))

}
