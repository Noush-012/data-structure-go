package main

import "fmt"

func binarySearch(ar []int, l, h, k int) int {

	if h >= l {
		m := l + (h-l)/2
		if ar[m] == k {
			return m
		}
		if ar[m] > k {
			return binarySearch(ar, l, m-1, k)
		} else {
			return binarySearch(ar, m+1, h, k)
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	n := len(arr)
	key := 10
	result := binarySearch(arr, 0, n-1, key)
	if result == -1 {
		fmt.Println("Element not found !!!")
	} else {
		fmt.Println("Value: ", arr[result], " found at index ", result)
	}
}
