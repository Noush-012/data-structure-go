package main

import "fmt"

func partition(arr []int, l int, h int) int {
	var pivot, i, j int
	pivot = arr[l]
	i = l
	j = h
	for i < j {
		for arr[i] <= pivot && i < h {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func quickSort(arr []int, l int, h int) {
	if l < h {
		p := partition(arr, l, h)
		quickSort(arr, l, p-1)
		quickSort(arr, p+1, h)
	}
}
func main() {
	// var array = []int{5, 8, -1, 6, 0, -8, -1, 0, 2, 3, 7, 58, 40, 60, -10, 2, 9, 74}
	array := []int{5, 6, 1, 8, -3, 0, 7, 95, 2}
	n := len(array)
	quickSort(array, 0, n-1)
	fmt.Println(array)

}
