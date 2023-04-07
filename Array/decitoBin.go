package main

import "fmt"

func deciToBin(dec int) {
	var decBin [5]int
	i := 0
	for dec > 0 {
		decBin[i] = dec % 2
		dec = dec / 2
		i++
	}
	for i := len(decBin) - 1; i >= 0; i-- {
		fmt.Println(decBin[i])
	}

}

func main() {
	dec := 25

	deciToBin(dec)
}
