package main

import "fmt"

var n = 1

func odd() {
	if n <= 10 {
		fmt.Println(n + 1)
		n++
	}
	even()

}
func even() {
	if n <= 10 {
		fmt.Println(n - 1)
		n++
	}
	odd()

}

func main() {

	odd()
}
