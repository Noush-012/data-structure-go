package main

import "fmt"

func main() {
	characters := []rune{'d', 'c', 'a', 'b'}
	n := len(characters)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if characters[j] > characters[j+1] {
				characters[j], characters[j+1] = characters[j+1], characters[j]
			}
		}
	}
	fmt.Println(string(characters))
}
