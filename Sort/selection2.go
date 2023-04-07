package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
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
	selectionSort(arr)
	fmt.Println("Sorted array is: ", arr)
}
