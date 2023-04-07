package main

import "fmt"

func printHelloWorld(n int) {
	if n == 0 {
		return
	}
	fmt.Println("Hello World")
	printHelloWorld(n - 1)
}

func main() {
	printHelloWorld(5)
}
