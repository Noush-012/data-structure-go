package main

import "fmt"

func missingElement(arr []int) {
	// n := len(arr)
	var result []int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != 1 && arr[i+1] != arr[i] {
			// fmt.Println(arr[i] + 1)

			for j := arr[i] + 1; j < arr[i+1]; j++ {
				result = append(result, j)
			}
		}
	}

	fmt.Println(result)

}

func main() {
	arr := []int{1, 3, 5, 5, 6, 7, 9}
	missingElement(arr)
}
