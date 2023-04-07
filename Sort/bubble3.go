package main

import (
	"fmt"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	fmt.Println("Enter the number of elements:")
	var n int
	fmt.Scan(&n)

	fmt.Println("Enter the elements:")
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Unsorted array is:", arr)
	bubbleSort(arr)
	fmt.Println("Sorted array is: ", arr)
}
