package main

import "fmt"

func BinToDeci(arr []int) int {
	n := len(arr)
	total := 0
	for i := 0; i < n; i++ {
		total = total*2 + arr[i]
	}
	return total

}
func main() {
	var arr = []int{0, 0, 1, 1, 0}

	result := BinToDeci(arr)
	fmt.Println(result)

}
