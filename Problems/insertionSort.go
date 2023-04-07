package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key:= arr[i]
		j:= i-1
		for j>=0 && arr[j] > key{
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
func main() {
	ar := []int{5, 8, -1, 6, 0, -8, -1, 0, 2, 3, 7, 58, 40, 60, -10, 2, 9, 74}
	insertionSort(ar)
	fmt.Println(ar)

}
