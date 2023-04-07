package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for arr[j] > key && j >= 0 {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key

	}
}
func main() {
	var n int

	fmt.Println("Enter array size:")
	fmt.Scanf("%d", &n)
	var ar [10]int
	for i := 0; i < n; i++ {
		fmt.Println("Enter element", i+1, ":")
		fmt.Scanf("%d", &ar[i])
	}

	// arr := []int{5, 8, -1, 6, 0, -8, -1, 0, 2, 3, 7, 58, 40, 60, -10, 2, 9, 74}
	//insertionSort(ar)
	fmt.Println(ar)

}
