package main

import "fmt"

func main() {
	var a, b, r, min int
	fmt.Println("Enter two integer:")
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a < b {
		min = a
	} else {
		min = b
	}
	for i := 1; i <= min; i++ {
		if a%i == 0 && b%i == 0 {
			r = i
		}

	}
	fmt.Println(r)
	// a<b
	// min = a

}
