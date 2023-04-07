package main

import "fmt"

func mergeSort(arr []int, l int, h int) {
	if l < h {

		m := (l + h) / 2
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, h)
		merge(arr, l, m, h)
	}
}
func merge(arr []int, l int, m int, h int) {

	var i, j, k int
	var b [18]int
	k = l 					// index for new array to store sorted
	i = l 					// left array start index 
	j = m + 1				 // right array start index
	for i <= m && j <= h { // 
		if arr[i] <= arr[j] {	 // logic for ascending sort
			b[k] = arr[i]  
			i++
		} else {
			b[k] = arr[j]
			j++
		}
		k++
	}
	if i > m {   // scenario if any of the sub array exhausted
		for j <= h { // for add right array balance
			b[k] = arr[j]
			j++
			k++
		}

	} else {
		for i <= m { // for add left array balance
			b[k] = arr[i]
			i++
			k++
		}
	}
	for k := l; k <= h; k++ {
		arr[k] = b[k]
	}

}

func main() {
	var arr = []int{5, 8, -1, 6, 0, -8, -1, 0, 2, 3, 7, 58, 40, 60, -10, 2, 9, 74}
	n := len(arr)
	mergeSort(arr, 0, n-1)
	fmt.Println(arr)

}
