package main

import (
	"fmt"
	"strings"
)

func countLetter(str string, letter string) int {
	return strings.Count(str, letter)
}

func main() {
	str := "Hey Greetings"
	letter := "s"
	fmt.Printf("The letter '%s' appears %d times in the string '%s'\n", letter, countLetter(str, letter), str)
}
