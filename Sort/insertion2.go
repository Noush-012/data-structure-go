package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
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
	insertionSort(arr)
	fmt.Println("Sorted array is: ", arr)
}
