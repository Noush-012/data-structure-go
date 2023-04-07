package main

import "fmt"

func removeAdjacent(str string) string {
	// Create a stack of runes to store the characters
	stack := []rune{}

	// Iterate over the characters in the input string
	for _, char := range str {
		// If the stack is not empty and the top character is the same as the current character, pop the top character
		if len(stack) > 0 && stack[len(stack)-1] == char {
			stack = stack[:len(stack)-1]
		} else {
			// Otherwise, push the current character onto the stack
			stack = append(stack, char)
		}
	}

	// Convert the stack back into a string and return it
	return string(stack)
}

func main() {
	// Test the function with some example inputs
	fmt.Println(removeAdjacent("abbaca")) // expected output: "ca"
	fmt.Println(removeAdjacent("aaaaaa")) // expected output: ""
	fmt.Println(removeAdjacent("abcde"))  // expected output: "abcde"
}
