package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}
	}

}
func main() {

	ar := []int{1, 9, 8, 7, 50, -2, 9, 8, 20, 65, -8, 0, 3}
	bubbleSort(ar)
	fmt.Println(ar)

}
