package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n-1; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]

		}

	}
}

func main() {

	ar := []int{5, 8, -1, 6, 0, -8, -1, 0, 2, 3, 7, 58, 40, 60, -10, 2, 9, 74}
	selectionSort(ar)
	fmt.Println(ar)

}
